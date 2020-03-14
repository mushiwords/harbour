package main

import (
	"github.com/astaxie/beego"
)

func main(){
	beego.StaticDir["/static"] = "static"
	beego.Run()
}
