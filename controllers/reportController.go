package controllers

import (
	"strconv"
	"vote_admin/models"
)

type ReportController struct {
	BaseController
}

func (this *ReportController) Get() {
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param("0"))

	quest:=models.GetQuestionInfo(question_id)
	question:=models.GetReport(quest)

	this.Data["question"] =question
	this.TplName = "admin/vote/record.html"
}

func (this *ReportController) FillList(){
	question_id,_ := strconv.Atoi( this.Input().Get("question_id"))
	selector_id,_ := strconv.Atoi( this.Input().Get("selector_id"))
	page,_ := strconv.Atoi( this.Input().Get("page"))
	limit:=5

	type JsonStruct struct{
		Code int
		Msg string
		Content *models.FillData
	}

	report:=models.FillReport(question_id,selector_id,limit,page)

	data:=&JsonStruct{0,"成功", report}

	this.Data["json"] = data
	this.ServeJSON()
	return
}