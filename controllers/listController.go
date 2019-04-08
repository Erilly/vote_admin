package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"strconv"
	"vote_admin/models"
)

type ListController struct {
	BaseController
}

func (this *ListController) Get() {
	page,_:= this.GetInt("page")
	title:= this.GetString("title")

	this.Data["questions"],this.Data["paginator"],_ = models.GetAllQuestion(page,title)
	this.Data["title"]=title

	this.TplName = "admin/vote/list.html"

}

func (this *ListController) Post() {

	//this.Ctx.WriteString( beego.AppConfig.String("mysqluser") + "test ini" )

	this.TplName = "admin/vote/list.html"
}

func (this *ListController) Create() {
	title:="问卷标题"
	description:="感谢您能抽出几分钟时间来参加本次答题，现在我们就马上开始吧！"

	question_id,_:=models.AddQuestion(title,description)

	this.Redirect("/list/edit/"+strconv.Itoa(int(question_id)),302)
	return
}

func (this *ListController) Edit(){
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param("0"))
	this.Data["question"] = models.GetQuestionInfo(question_id)
	this.TplName = "admin/vote/create.html"

}

func (this *ListController) Preview() {
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param("0"))
	this.Data["question"] = models.GetQuestionInfo(question_id)
	this.Data["isshow"] = false

	this.Layout = ""
	this.TplName = "index.html"

}

func (this *ListController) Publish() {
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param("0"))

	question:=models.GetQuestionInfo(question_id)
	question.PublishStatus=1

	_,err :=models.GetDatabase().C(models.MONGO_COLLECTION_QUESTION).Upsert(bson.M{"id":question_id},question)

	if err==nil{
		models.UpdateQuestion(&question)
	}

	url := "http://"+this.Ctx.Request.Host+"/questionpage-"+this.Ctx.Input.Param("0")

	qr := template.URL("data:image/png;base64,"+models.Qrcode(url))

	this.Data["questionpage"] =url
	this.Data["qr"] = qr

	this.Layout = ""
	this.TplName = "publish.html"
}

func (this *ListController) Report() {

}
