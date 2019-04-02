package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"vote_admin/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param(":id"))
	var questionInfo models.Question

	models.GetDatabase().C(models.MONGO_COLLECTION_QUESTION).Find(bson.M{"id":question_id}).One(&questionInfo)

	this.Data["question"] = questionInfo
	this.Data["isshow"] = true

	this.TplName = "index.html"
}

func (this *IndexController) Post(){
	ask:=this.Ctx.Input.RequestBody

fmt.Println(ask)
	var data *JSONStruct
	var content  string
	data=&JSONStruct{0,"更新失败", content}

	this.Data["json"] = data
	this.ServeJSON()
	return
}
