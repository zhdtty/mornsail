package main

import (
	"golog"
	//	"log"
	"net"
	"servlet"
	"sync"
	"time"
)

type ServerModule struct {
	//作为服务端
	//作为客户端
	mutex    sync.Mutex
	name     string
	address  string
	timeout  time.Duration
	listener net.Listener
	close    bool
	disp     *servlet.Dispatcher
	//	connPool map[string]*Connection
}

func NewServer(name string, address string, timeout time.Duration, disp *servlet.Dispatcher) *ServerModule {
	srv := &ServerModule{
		name:    name,
		address: address,
		timeout: timeout,
		close:   false,
		disp:    disp,
		//		connPool: make(map[string]*Connection),
	}
	return srv
}

func (srv *ServerModule) Start() {
	go srv.run()
}

func (srv *ServerModule) run() {
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

		golog.Debug("ServerModule", "run", "client connected", "remove.address", conn.RemoteAddr().String())

		go srv.GoSession(conn)
	}
}

func (srv *ServerModule) GoSession(conn net.Conn) {
	session := servlet.NewPlayerSession(conn, srv.timeout, srv.disp)
	/*
		srv.mutex.Lock()
		srv.connPool[conn.RemoteAddr().String()] = connection
		srv.mutex.Unlock()
	*/
	go session.Writer()
	go session.Reader()
}

func (srv *ServerModule) Close() {
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
