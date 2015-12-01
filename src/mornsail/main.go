package main

import (
	"config"
	"console"
	"flag"
	"fmt"
	//	"glog"
	"golog"
	//"path"
	"servlet"
	"time"
	"timer"
)

var FLAG_CONFIG_VAL = flag.String(
	"flag_val", "aaaaaaaaaaaaaaaaaa", "Config file to read flags from.")

func main() {
	//	flag.Parse()
	console.Init()
	config.Parse()

	srvModule := NewServer("Player Server", "10.0.253.61:6666", 10000000, servlet.G_dispatcher)
	srvModule.Start()

	st := timer.NewServerTimer()

	st.AddDailyTimer(19, 50, 0, func() {
		//		fmt.Println("daily timer")
	}, false)

	st.AddHourlyTimer(50, 0, func() {
		fmt.Println("hourly timer")
	}, false)

	st.AddIntervalTimer(5, func() {
		//		fmt.Println("interval timer")
	}, true)

	/*
		fmt.Println(*config.FLAG_CONFIG_FILE)
		fmt.Println(*config.FLAG_IP)
		fmt.Println("----------FLAG VAL : ", flag.Lookup("flag_val").Name)
		fmt.Println("----------FLAG VAL : ", flag.Lookup("flag_val").Value)
		fmt.Println("----------FLAG VAL : ", flag.Lookup("flag_val").Usage)
		fmt.Println("----------FLAG VAL : ", flag.Lookup("flag_val").DefValue)
	*/

	//	glog.Debug("test glog")

	//golog.SetLevel(golog.LevelDebug)
	golog.Info("main", "main", "Starting server", "ip", "10.0.253.61", "port", 6666)

	/*
		sysFilePath := path.Join("./", "sys.log")
		sysFile, err := golog.NewRotatingFileHandler(sysFilePath, 1024*1024, 10)
		//	sysFile, err := golog.NewTimeRotatingFileHandler(sysFilePath, golog.WhenSecond, 10)
		if err != nil {
			fmt.Printf("new log file error:%v\n", err.Error())
			return
		}

		golog.GlobalSysLogger = golog.New(sysFile, golog.Lfile|golog.Ltime|golog.Llevel)
		golog.GlobalSysLogger.SetLevel(golog.LevelDebug)
	*/
	/*
	   	sqlFilePath := path.Join(cfg.LogPath, sqlLogName)
	   	sqlFile, err := golog.NewRotatingFileHandler(sqlFilePath, MaxLogSize, 1)
	   	if err != nil {
	   		fmt.Printf("new log file error:%v\n", err.Error())
	                           return
	   	}
	   	golog.GlobalSqlLogger = golog.New(sqlFile, golog.Lfile|golog.Ltime|golog.Llevel)
	*/
	/*
		for i := 0; i < 10000000; i++ {
			//		fmt.Println(i)
			time.Sleep(time.Microsecond * 1000)
			golog.Debug("Module", "golog.Debug", "no any msg", uint32(i), "hello world")
		}
	*/
	time.Sleep(time.Second * 10000000)
}
