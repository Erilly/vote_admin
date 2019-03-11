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
	case SINGLE_SELECTOTR:
		this.TplName = "admin/templates/single-selector.html"
	case MULTI_SELECTOTR:
		this.TplName = "admin/templates/multi-selector.html"
	case SCORE_SELECTOTR:
		this.TplName = "admin/templates/score-selector.html"
	case SCORE_MATRIX_SELECTOTR:
		this.TplName = "admin/templates/score-matrix-selector.html"
	case FILL_SELECTOTR:
		this.TplName = "admin/templates/fill-selector.html"
	case FILL_MATRIX_SELECTOTR:
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

}


func (this *SelectorController) Edit(){


	this.Data["question"],_ = models.GetQuestionInfo(this.Ctx.Input.Param("0"))
	this.TplName = "admin/vote/create.html"

}