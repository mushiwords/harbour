package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"harbour/src/models"
	_ "harbour/src/models"
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
	this.Ctx.WriteString("hello")
}
/**
 * CaptainHandler
 */
func (this *CaptainHandler) Get() {
	pin := this.GetString("pin")
	fmt.Print("pin:",pin)
	info, err := models.SelectUserInfo(pin)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}

	body, err := json.Marshal(info)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	this.Ctx.WriteString(string(body))
}