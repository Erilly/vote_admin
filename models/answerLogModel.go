package models

import (
	"github.com/astaxie/beego/orm"
)

func AddAnswerLog(question_id int,content string) (int,error){
	o := orm.NewOrm()
	answer:=&AnswerLog{
		QuestionId:question_id,
		Content:content,
	}

	id,err:=o.Insert(answer)

	return int(id),err
}

func AnserLogCount(question_id int)(int){
	o := orm.NewOrm()
	count,_:=o.QueryTable(AnswerLog{}).Filter("QuestionId",question_id).Count()

	return int(count)
}