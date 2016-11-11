package logic

import (
	"config"
	"driver"
	"golog"
)

func Init() {
	golog.Info("Init", "Init", "-------init db driver")
	driver.Init()

	golog.Info("Init", "Init", "-------init player session manager")
	G_playerSessionMgr = NewPlayerSessionManager()

	//golog.Info("Init", "Init", "-------init socket dispatcher")
	//	G_dispatcher = NewDispatch()

	golog.Info("Init", "Init", "-------init socket server")
	G_sockServer = NewSocketServer("Socket Server", *config.FLAG_ADDR, 10000000)
}
