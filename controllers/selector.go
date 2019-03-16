package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"vote_admin/models"
)

type SelectorController struct {
	beego.Controller
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

func (this *SelectorController) UpdateSelector(){
	selector_id,_ := strconv.Atoi( this.Input().Get("selector_id"))

	var data *JSONStruct
	var content  string

	if num:=models.UpdateSelector(&models.Selector{Id:selector_id});num>0{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除成功", content}
	}else{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除失败", content}
	}
	this.Data["json"] = data
	this.ServeJSON()
	return

}

func (this *SelectorController) DeleteSelector(){
	selector_id,_ := strconv.Atoi( this.Input().Get("selector_id"))
	var data *JSONStruct
	var content  string

	if num:=models.DeleteSelector(&models.Selector{Id:selector_id});num>0{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除成功", content}
	}else{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除失败", content}
	}
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

func (this *SelectorController) DeleteOption(){
	option_id,_:= strconv.Atoi(this.Input().Get("option_id"))

	var data *JSONStruct
	var content  string

	if num:=models.UpdateOption(&models.Option{Id:option_id,Status:1});num>0{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除成功", content}
	}else{
		content = strconv.Itoa(int(num))
		data=&JSONStruct{0,"删除失败", content}
	}
	this.Data["json"] = data
	this.ServeJSON()
	return
}