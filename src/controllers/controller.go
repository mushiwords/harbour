package controllers

import (
"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type UserController struct {
	beego.Controller
}
/**
 * MainController
 */
func (this *MainController) Get() {
	this.Get()
}
func (this *MainController) Post() {
	this.Post()
}
func (this *MainController) Patch() {
	this.Patch()
}
func (this *MainController) Delete() {
	this.Delete()
}
/**
 * UserController
 */
func (this *UserController) Get() {
	this.Get()
}
func (this *UserController) Post() {
	this.Post()
}
func (this *UserController) Patch() {
	this.Patch()
}
func (this *UserController) Delete() {
	this.Delete()
}