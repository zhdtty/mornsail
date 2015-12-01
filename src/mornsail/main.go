package main

import (
	"config"
	"console"
	//"flag"
	//"fmt"
	//	"glog"
	"golog"
	//"path"
	"servlet"
	"time"
	//"timer"
)

func main() {
	//	flag.Parse()
	console.Init()
	config.Parse()

	srvModule := NewServer("Player Server", *config.FLAG_ADDR, 10000000, servlet.G_dispatcher)
	srvModule.Start()

	//golog.SetLevel(golog.LevelDebug)
	golog.Info("main", "main", "Starting server", "addr", *config.FLAG_ADDR)

	time.Sleep(time.Second * 10000000)
}
