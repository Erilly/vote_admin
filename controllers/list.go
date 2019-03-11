package controllers

import (
	"vote_admin/models"
)

type ListController struct {
	BaseController
}

func (this *ListController) Get() {

	this.Data["questions"],_ = models.GetAllQuestion()
	this.TplName = "admin/vote/list.html"

}

func (this *ListController) Post() {

	//this.Ctx.WriteString( beego.AppConfig.String("mysqluser") + "test ini" )

	this.TplName = "admin/vote/list.html"
}

func (this *ListController) Create() {
	title:="问卷标题"
	description:="感谢您能抽出几分钟时间来参加本次答题，现在我们就马上开始吧！"

	_,question_id:=models.AddQuestion(title,description)

	this.Redirect("/list/edit/"+question_id,302)
	return
}

func (this *ListController) Edit(){


	this.Data["question"],_ = models.GetQuestionInfo(this.Ctx.Input.Param("0"))
	this.TplName = "admin/vote/create.html"

}