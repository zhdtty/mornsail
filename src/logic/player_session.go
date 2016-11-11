package logic

import (
	//	"bytes"
	//	"encoding/binary"
	//	"errors"
	"fmt"
	"io"
	//	"log"
	"golog"
	"net"
	"protocol"
	//	"sync"
	"time"
)

type PlayerSession struct {
	//	mutex      sync.Mutex
	sync       chan bool
	conn       net.Conn
	WriteCache chan []byte
	timeout    time.Duration
	close      bool
	readBuf    []byte //Use as read buf cache

	Role *Role
	Id   int64
}

func NewPlayerSession(conn net.Conn, timeout time.Duration) *PlayerSession {
	ps := &PlayerSession{
		sync:       make(chan bool, 1),
		conn:       conn,
		WriteCache: make(chan []byte, 2048),
		timeout:    timeout,
		close:      false,
		readBuf:    make([]byte, protocol.PACKET_MAX_SIZE+12), //length + ptype + cmd
	}
	return ps
}

func (ps *PlayerSession) Writer() {
	for b := range ps.WriteCache {
		if ps.close == true {
			golog.Debug("PlayerSession", "Writer", "session closed")
			break
		}
		if b == nil {
			golog.Debug("PlayerSession", "Writer", "get bytes nil")
			break
		}
		_, err := ps.conn.Write(b)
		if err != nil {
			golog.Debug("PlayerSession", "Writer", "write error", "err", err)
			break
		}
	}
	ps.Disconnect()
}

func (ps *PlayerSession) Reader() {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "Reader", "throw error", "msg", msg)
			ps.close = true
		}
	}()
	for {
		//		golog.Debug("PlayerSession", "Reader", "Cycle reading", "cnt", cnt)
		if ps.close == true {
			golog.Info("PlayerSession", "Reader", "Session closed")
			break
		}

		data, err := ps.readPacket()
		if err != nil {
			if err == io.EOF {
				golog.Debug("PlayerSession", "Reader", "client disconnected",
					"err", err, "addr", ps.conn.RemoteAddr().String())
			} else {
				golog.Error("PlayerSession", "Reader", "Unable to read packet", "err", err)
			}
			ps.close = true
			return
		}
		//if encryption, do decryption
		decrData := ps.decrypt(data)

		if ok := ps.dispatch(decrData); !ok {
			ps.close = true
			return
		}
	}
}

func (ps *PlayerSession) Disconnect() {
	//	ps.mutex.Lock()
	//	defer ps.mutex.Unlock()
	ps.sync <- true
	defer func() { <-ps.sync }()
	ps.conn.Close()
	golog.Debug("PlayerSession", "Writer", "conn closed")
	close(ps.WriteCache)
	golog.Debug("PlayerSession", "Writer", "write cache closed")
}

/*
func (session *PlayerSession) readPacket() ([]byte, error) {
	len, err := session.conn.Read(session.readBuf[0:]) //在协议包积压的情况下，这种读法会出现丢包
	if err != nil {
		return nil, err
	}
	if len < 4 {
		return nil, fmt.Errorf("Packet error, header invalid, length = %d", len)
	}
	header := session.readBuf[0:4]

	length := int(uint32(header[3]) | uint32(header[2])<<8 | uint32(header[1])<<16 | uint32(header[0]<<32))
	golog.Debug("PlayerSession", "readPacket", "packet length", "length", length)
	if length < 1 {
		return nil, fmt.Errorf("Packet error, length invalid, length = %d", length)
	}

	if length > protocol.PACKET_MAX_SIZE+8 { //ptype + cmd
		return nil, fmt.Errorf("Packet error, size = %d", length)
	}

	data := session.readBuf[4 : length+4]
	return data, nil
}
*/

func (ps *PlayerSession) readPacket() ([]byte, error) {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "readerPacket", "throw error", "msg", msg)
			panic(msg)
		}
	}()

	header := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(ps.conn, header); err != nil {
		return nil, err
		//		return nil, errors.New("Packet error, header invalid")
	}

	//parse header
	length := int(uint32(header[3]) | uint32(header[2])<<8 | uint32(header[1])<<16 | uint32(header[0]<<32))
	if length < 1 {
		return nil, fmt.Errorf("Packet error, length invalid, length = %d", length)
	}

	if length > protocol.PACKET_MAX_SIZE {
		return nil, fmt.Errorf("Packet error, size = %d", length)
	}

	data := make([]byte, length)
	if _, err := io.ReadFull(ps.conn, data); err != nil {
		return nil, fmt.Errorf("Packet error, invalid data, size = %d", length)
	}
	return data, nil
}

func (ps *PlayerSession) encrypt(data []byte) []byte {
	return data
}

func (ps *PlayerSession) decrypt(data []byte) []byte {
	return data
}

func (ps *PlayerSession) dispatch(data []byte) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "dispatch", "throw error", "msg", msg)
			panic(msg)
		}
	}()

	pack := protocol.ReceiveData(data)
	if pack == nil {
		golog.Error("PlayerSession", "dispatch", "Receive no data")
		return false
	}

	return G_dispatcher.Dispatch(ps, pack)
}

func (ps *PlayerSession) SendData(b []byte) {
	//	session.mutex.Lock()
	//	defer session.mutex.Unlock()
	if b == nil || ps.close == true {
		golog.Error("PlayerSession", "SendData", "Data is nil or session closed")
		return
	}
	/*
		if len(session.WriteCache) != cap(session.WriteCache) {
			golog.Error("PlayerSession", "SendData", "Channel full, close conn")
			session.Disconnect()
			return
		}
	*/

	ps.WriteCache <- b
}

func (ps *PlayerSession) SetRole(role *Role) {
	ps.Role = role
}
