package main

import (
	"config"
	"console"
	//"flag"
	//	"glog"
	"golog"
	//"path"
	"os"
	"os/signal"
	"servlet"
	"syscall"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)
	//	signal.Notify(c, os.Interrupt, os.Kill)

	sig := <-c
	golog.Info("main", "main", "Server begin close", "sig", sig)

	golog.Info("main", "main", "Closing server module ...")
	srvModule.Close()

	golog.Info("main", "main", "Server close success")
	time.Sleep(time.Second)
}
