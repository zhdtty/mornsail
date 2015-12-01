package servlet

import (
	"golog"
	"protocol"
)

type IServlet interface {
	DoRequest(session *PlayerSession, pack *protocol.Packet) bool
}

type Dispatcher struct {
	register map[int32]IServlet
}

func NewDispatch() *Dispatcher {
	disp := &Dispatcher{
		register: make(map[int32]IServlet),
	}
	return disp
}

func (disp *Dispatcher) Register(cmd int32, servlet IServlet) {
	disp.register[cmd] = servlet
}

func (disp *Dispatcher) Dispatch(session *PlayerSession, pack *protocol.Packet) bool {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("Dispatcher", "Dispatch", "Dispatch panic", "msg", msg)
			panic(msg)
		}
	}()
	if servlet, ok := disp.register[pack.Cmd]; ok {
		return servlet.DoRequest(session, pack)
	}
	return false
}

var G_dispatcher *Dispatcher = NewDispatch()
