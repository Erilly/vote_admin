package models

import (
	"time"
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

func GetQuestionInfo(questionID string)(*Question,error){
	o := orm.NewOrm()
	questions := new(Question)
	qs := o.QueryTable("vt_question")
	err:=qs.Filter("question_id",questionID).One(questions)
	return questions,err
}

func AddQuestion(title,content string) (int64,string){
	now:=time.Now()
	token:=getMd5Token(strconv.FormatInt(time.Now().UTC().UnixNano(), 10),8)

	o := orm.NewOrm()

	question:=&Question{
		Title:title,
		Description:content,
		QuestionId:token,
		Ctime:now,
		Mtime:now,
	}

	id,err:=o.Insert(question)
	if err==nil{

	}
	return id,token
}