package main

import (
	"flag"
	"config"
	"fmt"
	"os"
	"runtime"
	"common/mylog"
)

func main() {

	/**
	 * Run it: ./harbour -c etc/harbour.json
	 **/
	wordPtr := flag.String("c", "./etc/harbour.json", "Harbour configuration file.")
    flag.Parse()

	/** Load Configure **/
	config, err := config.LoadConfig(*wordPtr)
	if err != nil {
		fmt.Println("main", "load config failed: ", err.Error())
		os.Exit(-1)
		return
	}

    /** Init Log Service **/
	if err := mylog.Init(config.MyLog.AccLogFile,config.MyLog.InfoLogFile,
		config.MyLog.LogLevel, config.MyLog.LogMaxAge); err != nil {
		fmt.Println("Init Log faild: ", err.Error())
		os.Exit(-1)
		return
	}
	mylog.LogInfo("Harbour Server started.")

	/** Set CpuNumber **/
	runtime.GOMAXPROCS(runtime.NumCPU())
	mylog.LogInfo("Using %v cpus.", runtime.NumCPU())

	/** Start Service **/
	if err := Start(config.Service); err != nil {
		mylog.LogError("start service failed: %v", err)
		os.Exit(-1)
		return
	}
}
