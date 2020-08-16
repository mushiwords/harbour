package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"html/template"
	"monitor/src/worker"
)

type MainController struct {
	beego.Controller
}

type CaptainHandler struct {
	beego.Controller
}
/**
 * MainController
 */
func (this *MainController) Get() {
	t, err := template.ParseFiles("./src/views/index.html")
	if err != nil {
		this.Ctx.WriteString("err")
		return
	}
	t.Execute(this.Controller.Ctx.ResponseWriter,this.Controller.Ctx.Request)
	//this.Ctx.WriteString("hello")
}
/**
 * CaptainHandler
 */
func (this *CaptainHandler) Get() {
	//pin := this.GetString("pin")
	//fmt.Print("pin:",pin)
	//info, err := models.SelectUserInfo(pin)
	//if err != nil {
	//	this.Ctx.WriteString(err.Error())
	//}
//
	//body, err := json.Marshal(info)
	//if err != nil {
	//	this.Ctx.WriteString(err.Error())
	//}
	data,err := worker.ReadTxtFile("./src/static/load/xyz.txt")
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	body, err := json.Marshal(data)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	this.Ctx.WriteString(string(body))
}