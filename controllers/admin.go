package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) Get() {
	//this.Ctx.WriteString( beego.AppConfig.String("mysqluser") + "test ini" )

	this.Data["isVote"] = true
	this.Data["isVoteList"] = true

	this.Layout = "admin/layout/main.html"
	this.TplName = "admin/vote/list.html"
}
