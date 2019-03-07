package controllers

import (
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
}

//func (this *BaseController) Prepare() {
//	tu:=this.GetSession(SESSION_USER_KEY)
//
//	if tu==nil{
//		this.Redirect("/login",301)
//		return
//	}
//}