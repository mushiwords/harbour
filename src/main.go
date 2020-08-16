package main

import (
	"fmt"
	"github.com/astaxie/beego"
	config "github.com/astaxie/beego/config"
	_ "monitor/src/routers"
)

func main(){
	conf, err := config.NewConfig("ini", "src/config/app.conf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	beego.BConfig.AppName = "harbour"
	port, err := conf.Int("http::httpport")
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	beego.BConfig.Listen.HTTPPort = port
	beego.Run()
}
