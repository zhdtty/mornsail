package main

import (
	"logic"
)

type ServerModule struct {
}

func NewServer() *ServerModule {
	logic.Init()
	sm := &ServerModule{}
	return sm
}

func (srv *ServerModule) Start() {
	go logic.G_sockServer.Run()
}

func (srv *ServerModule) Close() {
	logic.G_sockServer.Close()
}
