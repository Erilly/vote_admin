package controllers

import (
	"vote_admin/models"
	"github.com/astaxie/beego"
	"strconv"
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
	_type,_ := strconv.Atoi( this.Input().Get("type"))

	this.TplName = "admin/templates/single-selector.html"

	switch _type {
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
	this.Data["type"] = _type
	single_temp,_:=this.RenderString()

	data:=&JSONStruct{0,strconv.Itoa(_type), single_temp}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *SelectorController) Post() {
	template_type,_ := strconv.Atoi( this.Input().Get("template_type"))

	this.TplName = "admin/templates/single-selector.html"
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

	models.AddSelector(title,this.Input().Get("question_id"),int8(template_type))

	this.Data["type"] = template_type
	single_temp,_:=this.RenderString()

	data:=&JSONStruct{0,strconv.Itoa(template_type), single_temp}
	this.Data["json"] = data
	this.ServeJSON()
	return
}


func (this *SelectorController) Edit(){


	this.Data["question"],_ = models.GetQuestionInfo(this.Ctx.Input.Param("0"))
	this.TplName = "admin/vote/create.html"

}