package models

import (
	"github.com/astaxie/beego/orm"
)


func GetAllQuestion()([]*Question,error){
	o := orm.NewOrm()
	questions := make([]*Question,0)
	qs := o.QueryTable(Question{})
	_,err:=qs.OrderBy("-ctime").All(&questions)

	return questions,err
}

func GetQuestion(question_id int)(Question,error){
	o := orm.NewOrm()
	question := Question{Id: question_id}
	err:=o.Read(&question)

	return question,err
}

func GetQuestionInfo(question_id int)(Question){
	o := orm.NewOrm()
	question := Question{Id: question_id}
	o.Read(&question)
	o.QueryTable(Selector{}).Filter("Question__id", question.Id).All(&question.Selector)

	for k,v:=range question.Selector{
		o.QueryTable(Option{}).Filter("Selector__id", v.Id).All(&question.Selector[k].Option)
	}
	return question
}

func AddQuestion(title,content string) (int,error){
	o := orm.NewOrm()
	question:=&Question{
		Title:title,
		Description:content,
		PageCut:1,
	}

	id,err:=o.Insert(question)

	return int(id),err
}


func UpdateQuestion(updateData *Question) (int64){
	o := orm.NewOrm()
	num,_ :=o.Update(updateData)
	return num
}
