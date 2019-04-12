package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"vote_admin/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	question_id:=this.Ctx.Input.Param(":id")
	var cv bool

	ap:=this.Ctx.GetCookie("author-page")

	for  _,val := range strings.Split(ap,"|"){
		if question_id == val{
			cv = true
		}
	}
	if cv==false {
		qid, _ := strconv.Atoi(question_id)

		qs:=models.GetDatabase().C(models.MONGO_COLLECTION_QUESTION).Find(bson.M{"id": qid})

		if num,_:=qs.Count();num<1{
			this.Data["warning"] = "该问卷不存在！"
			this.TplName = "warning.html"
		}else{
			var questionInfo models.Question
			qs.One(&questionInfo)

			this.Data["question"] = questionInfo
			this.Data["isshow"] = true

			this.TplName = "index.html"
		}

	}else{
		this.Data["warning"] = "您已提交该问卷，感谢您的参与！"
		this.TplName = "warning.html"
	}
}

func (this *IndexController) Post(){
	question_id:=this.Ctx.Input.Param(":id")
	answer:= this.Input().Get("answer")
	fmt.Println(question_id,answer)

	var data *JSONStruct
	var content  string
	var cv bool

	ap:=this.Ctx.GetCookie("author-page")

	for  _,val := range strings.Split(ap,"|"){
		if question_id == val{
			cv = true
		}
	}

	if cv==false{
		var f interface{}
		json.Unmarshal([]byte(answer), &f)
		qid,_:=strconv.Atoi(question_id)

		models.AddAnswerLog(qid, answer)
		models.GetDatabase().C(models.MONGO_COLLECTION_VOTE_LOG).Insert(f)

		this.Ctx.SetCookie("author-page", ap+"|"+question_id)
		data=&JSONStruct{0,"已成功提交", content}
	}else{
		data=&JSONStruct{1,"您已提交该问卷，感谢您的参与！", content}
	}

	this.Data["json"] = data
	this.ServeJSON()
	return
}
