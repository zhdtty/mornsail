package logic

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"golog"
	"io"
	"net"
	"protocol"
	"sync"
	"time"
)

type SocketServer struct {
	//作为服务端
	//作为客户端
	mutex    sync.Mutex
	name     string
	address  string
	timeout  time.Duration
	listener net.Listener
	close    bool
	//	connPool map[string]*Connection
}

func NewSocketServer(name string, address string, timeout time.Duration) *SocketServer {
	srv := &SocketServer{
		name:    name,
		address: address,
		timeout: timeout,
		close:   false,
		//		connPool: make(map[string]*Connection),
	}
	return srv
}

func (srv *SocketServer) Run() {
	var err error
	srv.listener, err = net.Listen("tcp", srv.address)
	if err != nil {
		golog.Fatal("ServerModule", "run", "Net listen err", "err", err)
	}
	if srv.listener == nil {
		golog.Fatal("ServerModule", "run", "Net listen nil")
	}

	golog.Debug("ServerModule", "run", "listen success", "servername", srv.name, "address", srv.address)

	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			return
		}

		golog.Debug("ServerModule", "run", "client connected", "address", conn.RemoteAddr().String())

		go srv.GoSession(conn)
	}
}

func (srv *SocketServer) GoSession(conn net.Conn) {
	session := NewPlayerSession(conn, srv.timeout)
	/*
		srv.mutex.Lock()
		srv.connPool[conn.RemoteAddr().String()] = connection
		srv.mutex.Unlock()
	*/
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("PlayerSession", "readerPacket", "throw error", "msg", msg)
			panic(msg)
		}
	}()

	header := make([]byte, 4)
	_, err := io.ReadFull(conn, header[0:4])
	if err != nil {
		golog.Debug("ServerModule", "GoSession", "conn close! read header failed", "address", conn.RemoteAddr().String())
		conn.Close()
		return
	}
	var flag int32
	buf := bytes.NewBuffer(header)
	binary.Read(buf, binary.LittleEndian, &flag)

	golog.Debug("ServerModule", "GoSession", "connected!", "flag", flag)
	gateState := &protocol.S2C_GateState{
		State: proto.Int32(0),
		Key:   proto.Int64(0),
	}
	data, err := proto.Marshal(gateState)
	if err != nil {
		golog.Debug("ServerModule", "GoSession", "conn close! Marshal gate state failed", "address", conn.RemoteAddr().String())
		conn.Close()
		return
	}
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, len(data)+4)
	buffer.Write([]byte(data))
	conn.Write(buffer.Bytes())

	golog.Debug("ServerModule", "GoSession", "Create player session")
	go session.Writer()
	go session.Reader()
}

func (srv *SocketServer) Close() {
	srv.close = true
	srv.listener.Close()
	/*
		srv.mutex.Lock()
		for conn, _ := range srv.connPool {
			conn.Close()
		}
		srv.connPool = make(map[string]*Connection)
		srv.mutex.Unlock()
	*/
}

var G_sockServer *SocketServer
