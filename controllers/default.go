package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString( beego.AppConfig.String("mysqluser") + "test ini" )
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplName = "index.tpl"
}
