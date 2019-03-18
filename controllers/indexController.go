package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["logininfo"] = this.GetSession(SESSION_USER_KEY)

	this.TplName = "index.tpl"
}
