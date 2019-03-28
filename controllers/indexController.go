package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"vote_admin/models"
)

type IndexController struct {
	beego.Controller
}

type DeploymentText struct{
	Title string
	Content string
}

func (this *IndexController) Get() {
	question_id,_:=strconv.Atoi(this.Ctx.Input.Param(":id"))
	var questionInfo models.Question

	models.GetDatabase().C(models.MONGO_COLLECTION_QUESTION).Find(bson.M{"id":question_id}).One(&questionInfo)

	this.Data["question"] = questionInfo

	this.TplName = "index.html"
}
