package controllers

import "github.com/astaxie/beego"

const (
	SESSION_USER_KEY = "SESSION_USER_KEY"

	//题型
	SINGLE_SELECTOTR = 1
	MULTI_SELECTOTR = 2
	SCORE_SELECTOTR = 3
	SCORE_MATRIX_SELECTOTR = 4
	FILL_SELECTOTR = 5
	FILL_MATRIX_SELECTOTR = 6
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare(){
	logininfo:=this.GetSession(SESSION_USER_KEY)

	if logininfo==nil{
		this.Redirect("/login",302)
	}else{
		this.Data["logininfo"] = logininfo
	}

	this.Data["isVote"] = true
	this.Data["isVoteList"] = true
	this.Data["logininfo"] = this.GetSession(SESSION_USER_KEY)
	this.Layout = "admin/layout/main.html"
}