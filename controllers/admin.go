package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	//this.Ctx.WriteString( beego.AppConfig.String("mysqluser") + "test ini" )
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"

	this.Data["isVote"] = true
	this.Data["isVoteList"] = true

	this.Layout = "admin/layout/main.html"
	this.TplName = "admin/vote/list.html"
}
