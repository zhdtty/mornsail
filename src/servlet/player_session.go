package servlet

import (
	//	"bytes"
	//	"encoding/binary"
	//	"errors"
	"fmt"
	"io"
	//	"log"
	"golog"
	"net"
	"player"
	"protocol"
	"sync"
	"time"
)

type PlayerSession struct {
	mutex      sync.Mutex
	conn       net.Conn
	WriteCache chan []byte
	timeout    time.Duration
	disp       *Dispatcher
	close      bool
	readBuf    []byte //Use as read buf cache

	//custom
	Role *player.Role
	Id   int64
}

func NewPlayerSession(conn net.Conn, timeout time.Duration, disp *Dispatcher) *PlayerSession {
	session := &PlayerSession{
		conn:       conn,
		WriteCache: make(chan []byte, 2048),
		disp:       disp,
		timeout:    timeout,
		close:      false,
		readBuf:    make([]byte, protocol.PACKET_MAX_SIZE+12), //length + ptype + cmd
	}
	return session
}

func (session *PlayerSession) Writer() {
	for b := range session.WriteCache {
		if session.close == true {
			golog.Debug("PlayerSession", "Writer", "session closed")
			break
		}
		if b == nil {
			golog.Debug("PlayerSession", "Writer", "get bytes nil")
			break
		}
		_, err := session.conn.Write(b)
		if err != nil {
			golog.Debug("PlayerSession", "Writer", "write error", "err", err)
			break
		}
	}
	session.Disconnect()
}

func (session *PlayerSession) Reader() {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "Reader", "throw error", "msg", msg)
			session.close = true
		}
	}()
	for {
		//		golog.Debug("PlayerSession", "Reader", "Cycle reading", "cnt", cnt)
		if session.close == true {
			golog.Info("PlayerSession", "Reader", "Session closed")
			break
		}

		data, err := session.readPacket()
		if err != nil {
			if err == io.EOF {
				golog.Debug("PlayerSession", "Reader", "client disconnected",
					"err", err, "addr", session.conn.RemoteAddr().String())
			} else {
				golog.Error("PlayerSession", "Reader", "Unable to read packet", "err", err)
			}
			session.close = true
			return
		}
		//if encryption, do decryption
		decrData := session.decrypt(data)

		if ok := session.dispatch(decrData); !ok {
			session.close = true
			return
		}
	}
}

func (session *PlayerSession) Disconnect() {
	session.mutex.Lock()
	defer session.mutex.Unlock()
	session.conn.Close()
	close(session.WriteCache)
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

func (session *PlayerSession) readPacket() ([]byte, error) {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "readerPacket", "throw error", "msg", msg)
			panic(msg)
		}
	}()

	header := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(session.conn, header); err != nil {
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
	if _, err := io.ReadFull(session.conn, data); err != nil {
		return nil, fmt.Errorf("Packet error, invalid data, size = %d", length)
	}
	return data, nil
}

func (session *PlayerSession) encrypt(data []byte) []byte {
	return data
}

func (session *PlayerSession) decrypt(data []byte) []byte {
	return data
}

func (session *PlayerSession) dispatch(data []byte) bool {
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

	return session.disp.Dispatch(session, pack)
}

func (session *PlayerSession) SendData(b []byte) {
	//	session.mutex.Lock()
	//	defer session.mutex.Unlock()
	if b == nil || session.close == true {
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

	session.WriteCache <- b
}

func (session *PlayerSession) SetRole(role *player.Role) {
	session.Role = role
}
