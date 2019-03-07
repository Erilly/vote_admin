package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "admin/vote/login.html"
}

func (this *LoginController) Post() {

	username:=this.Input().Get("username")
	password:=this.Input().Get("password")

	if beego.AppConfig.String("adminuser") == username &&
		beego.AppConfig.String("adminpass") == password{
		this.Ctx.SetCookie("username",username,0)
		this.Ctx.SetCookie("password",password,0)
	}
	this.Redirect("/",301)
	return
}