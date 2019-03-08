package controllers

import (
	"vote_admin/models"
	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
}

func (this *ListController) Prepare(){
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

	//this.Data["question"],_ = models.GetQuestion(this.Ctx.Input.Param("0"))
	this.Data["question"],_ = models.GetQuestion("1b9fdeb7")
	this.TplName = "admin/vote/create.html"

}