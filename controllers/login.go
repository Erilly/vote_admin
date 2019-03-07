package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit:= this.Input().Get("exit") == "true"
	if isExit{
		this.DelSession(SESSION_USER_KEY)
	}
	this.TplName = "admin/vote/login.html"
}

func (this *LoginController) Post() {

	username:=this.Input().Get("username")
	password:=this.Input().Get("password")

	if beego.AppConfig.String("adminuser") == username &&
		beego.AppConfig.String("adminpass") == password{

		this.SetSession(SESSION_USER_KEY,username)
		this.Redirect("/admin",301)
	}
	this.TplName = "admin/vote/login.html"
}