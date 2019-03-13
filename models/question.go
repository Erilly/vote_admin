package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)


func GetAllQuestion()([]*Question,error){
	o := orm.NewOrm()
	questions := make([]*Question,0)
	qs := o.QueryTable("vt_question")
	_,err:=qs.OrderBy("-ctime").All(&questions)

	return questions,err
}

func GetQuestionInfo(questionID string)(Question){
	o := orm.NewOrm()
	qid,_:= strconv.Atoi(questionID)
	questions := Question{Id: uint(qid)}
	o.Read(&questions)

	return questions
}

func AddQuestion(title,content string) (uint,error){
	o := orm.NewOrm()
	question:=&Question{
		Title:title,
		Description:content,
		PageCut:1,
	}

	id,err:=o.Insert(question)

	return uint(id),err
}