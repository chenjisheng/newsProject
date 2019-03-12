package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*
网站icon
uri: /favicon.ico
method: get
 */
func (this *MainController) Get() {
	this.Redirect("static/new.icon",302)
}
