package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Prepare(){
	logininfo:=this.GetSession(SESSION_USER_KEY)

	if logininfo==nil{
		this.Redirect("/login",302)
	}else{
		this.Data["logininfo"] = logininfo
	}
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["logininfo"] = this.GetSession(SESSION_USER_KEY)

	this.TplName = "index.tpl"
}
