package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


func GetAllSelector()([]*Question,error){
	o := orm.NewOrm()
	questions := make([]*Question,0)
	qs := o.QueryTable("vt_question")
	_,err:=qs.OrderBy("-ctime").All(&questions)

	return questions,err
}

func GetSelectorInfo(questionID string)(*Question,error){
	o := orm.NewOrm()
	questions := new(Question)
	qs := o.QueryTable("vt_question")
	err:=qs.Filter("question_id",questionID).One(questions)
	return questions,err
}

func AddSelector(title,question_id string, template_type int8) (string){
	selector_id:=getMd5Token(16,"")

	o := orm.NewOrm()

	question:=&Selector{
		Title:title,
		QuestionId:question_id,
		SelectorId:selector_id,
		TemplateType:template_type,
		Ctime:time.Now(),
		Mtime:time.Now(),
	}

	_,err:=o.Insert(question)

	if err == nil {
		options := selectorTemplate(selector_id,template_type)
		o.InsertMulti(len(options),options)
	}
	return selector_id
}

func selectorTemplate(selector_id string,template_type int8)([]Option){
	var multiOptions []Option
	switch template_type{
	case SINGLE_SELECTOTR:
		multiOptions= []Option{
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
		}
	case MULTI_SELECTOTR:
		multiOptions= []Option{
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
		}
	case SCORE_SELECTOTR:
		multiOptions= []Option{
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
		}
	case SCORE_MATRIX_SELECTOTR:
	case FILL_SELECTOTR:
		multiOptions= []Option{
			Option{
				SelectorId: selector_id,
				OptionId:   getMd5Token(16, ""),
				Title:      "选项标题",
				Ctime:      time.Now(),
				Mtime:      time.Now(),
			},
		}
	case FILL_MATRIX_SELECTOTR:
	}


	return multiOptions
}