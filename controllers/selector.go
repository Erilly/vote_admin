package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"vote_admin/models"
)


type JSONStruct struct {
	Code int
	Msg  string
	Content string
}

type SelectorController struct {
	beego.Controller
}

func (this *SelectorController) Prepare(){

}

func (this *SelectorController) Get() {
	template_type,_ := strconv.Atoi( this.Input().Get("template_type"))
	question_id,_ := strconv.Atoi( this.Input().Get("question_id"))

	switch template_type {
	case models.SINGLE_SELECTOTR:
		this.TplName = "admin/templates/single-selector.html"
	case models.MULTI_SELECTOTR:
		this.TplName = "admin/templates/multi-selector.html"
	case models.SCORE_SELECTOTR:
		this.TplName = "admin/templates/score-selector.html"
	case models.SCORE_MATRIX_SELECTOTR:
		this.TplName = "admin/templates/score-matrix-selector.html"
	case models.FILL_SELECTOTR:
		this.TplName = "admin/templates/fill-selector.html"
	case models.FILL_MATRIX_SELECTOTR:
		this.TplName = "admin/templates/fill-matrix-selector.html"
	}
	this.Data["type"] = template_type
	options := models.GetSelector(question_id)
	this.Data["selector"] = options

	single_temp,_:=this.RenderString()

	data:=&JSONStruct{0,strconv.Itoa(template_type), single_temp}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *SelectorController) Post() {
	template_type,_ := strconv.Atoi( this.Input().Get("template_type"))
	question_id,_ := strconv.Atoi(this.Input().Get("question_id"))

	var title string

	switch template_type {
		case models.SINGLE_SELECTOTR:
			this.TplName = "admin/templates/single-selector.html"

			title="单项选择"
		case models.MULTI_SELECTOTR:
			this.TplName = "admin/templates/multi-selector.html"

			title="多项选择"
		case models.SCORE_SELECTOTR:
			this.TplName = "admin/templates/score-selector.html"
			title="打分"
		case models.SCORE_MATRIX_SELECTOTR:
			this.TplName = "admin/templates/score-matrix-selector.html"
			title="矩阵打分"
		case models.FILL_SELECTOTR:
			this.TplName = "admin/templates/fill-selector.html"
			title="填空"

		case models.FILL_MATRIX_SELECTOTR:
			this.TplName = "admin/templates/fill-matrix-selector.html"
			title="矩阵填空"
	}



	this.Data["selector"] = models.AddSelector(title,question_id,template_type)
	single_temp,_:=this.RenderString()

	data:=&JSONStruct{0,strconv.Itoa(template_type), single_temp}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *SelectorController) AddOption() {
	template_type,_ := strconv.Atoi( this.Input().Get("template_type"))
	selector_id,_ := strconv.Atoi( this.Input().Get("selector_id"))

	switch template_type {
	case models.SINGLE_SELECTOTR:
		this.TplName = "admin/templates/single-option.html"

	case models.MULTI_SELECTOTR:
		this.TplName = "admin/templates/multi-option.html"

	case models.SCORE_SELECTOTR:
		this.TplName = "admin/templates/score-option.html"
	case models.SCORE_MATRIX_SELECTOTR:
		this.TplName = "admin/templates/score-matrix-option.html"
	case models.FILL_SELECTOTR:
		this.TplName = "admin/templates/fill-option.html"

	case models.FILL_MATRIX_SELECTOTR:
		this.TplName = "admin/templates/fill-matrix-option.html"
	}

	options := models.AddOption(selector_id,template_type)

	this.Data["type"] = template_type
	this.Data["option"] = options
	single_temp,_:=this.RenderString()

	data:=&JSONStruct{0,strconv.Itoa(template_type), single_temp}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *SelectorController) Edit(){
	question_id,_:= strconv.Atoi(this.Ctx.Input.Param("0"))

	this.Data["question"] = models.GetQuestionInfo(question_id)
	this.TplName = "admin/vote/create.html"

}