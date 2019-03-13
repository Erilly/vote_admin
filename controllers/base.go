package controllers

import "github.com/astaxie/beego"

const (
	SESSION_USER_KEY = "SESSION_USER_KEY"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare(){
	logininfo:=this.GetSession(SESSION_USER_KEY)

	if logininfo==nil{
		//this.Redirect("/login",302)
	}else{
		this.Data["logininfo"] = logininfo
	}

	this.Data["isVote"] = true
	this.Data["isVoteList"] = true
	this.Data["logininfo"] = this.GetSession(SESSION_USER_KEY)
	this.Layout = "admin/layout/main.html"
}