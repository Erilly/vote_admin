package controllers

import (
	"fmt"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["username"] = fmt.Sprint(this.GetSession(SESSION_USER_KEY))

	this.TplName = "index.tpl"
}
