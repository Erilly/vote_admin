package models
//
//import (
//	"github.com/astaxie/beego/orm"
//	"time"
//)
//
//
//func GetAllSelector()([]*Question,error){
//	o := orm.NewOrm()
//	questions := make([]*Question,0)
//	qs := o.QueryTable("vt_question")
//	_,err:=qs.OrderBy("-ctime").All(&questions)
//
//	return questions,err
//}
//
//func GetSelectorInfo(questionID string)(*Question,error){
//	o := orm.NewOrm()
//	questions := new(Question)
//	qs := o.QueryTable("vt_question")
//	err:=qs.Filter("question_id",questionID).One(questions)
//	return questions,err
//}
//
//func AddSelector(title,question_id int, template_type int8) (string){
//	selector_id:=getMd5Token(16,"")
//
//	o := orm.NewOrm()
//
//	question:=&Selector{
//		Title:title,
//		QuestionId:question_id,
//		TemplateType:template_type,
//		Ctime:time.Now(),
//		Mtime:time.Now(),
//	}
//
//	_,err:=o.Insert(question)
//
//	if err == nil {
//		options := selectorTemplate(selector_id,template_type,2)
//		o.InsertMulti(len(options),options)
//	}
//	return selector_id
//}
//
//func AddOption(question_id string, template_type int8) ([]Option){
//	selector_id:=getMd5Token(16,"")
//	options := []Option{}
//	o := orm.NewOrm()
//	options = selectorTemplate(selector_id,template_type,2)
//	o.InsertMulti(len(options),options)
//
//	return options
//}
//
//func selectorTemplate(selector_id string,template_type int8,createNum int)([]Option){
//	if createNum == 0 {
//		createNum = 1
//	}
//	multiOptions :=make([]Option,createNum)
//
//	switch template_type{
//	case SINGLE_SELECTOTR:
//		for i:=0;i<createNum;i++{
//			opt := Option{
//				SelectorId: selector_id,
//				OptionId:   getMd5Token(16, ""),
//				Title:      "选项标题",
//				Ctime:      time.Now(),
//				Mtime:      time.Now(),
//			}
//			multiOptions = append(multiOptions,opt)
//		}
//	case MULTI_SELECTOTR:
//		for i:=0;i<createNum;i++{
//			opt := Option{
//				SelectorId: selector_id,
//				OptionId:   getMd5Token(16, ""),
//				Title:      "选项标题",
//				Ctime:      time.Now(),
//				Mtime:      time.Now(),
//			}
//			multiOptions = append(multiOptions,opt)
//		}
//	case SCORE_SELECTOTR:
//		multiOptions= []Option{
//			Option{
//				SelectorId: selector_id,
//				OptionId:   getMd5Token(16, ""),
//				Title:      "选项标题",
//				Ctime:      time.Now(),
//				Mtime:      time.Now(),
//			},
//		}
//	case SCORE_MATRIX_SELECTOTR:
//	case FILL_SELECTOTR:
//		multiOptions= []Option{
//			Option{
//				SelectorId: selector_id,
//				OptionId:   getMd5Token(16, ""),
//				Title:      "选项标题",
//				Ctime:      time.Now(),
//				Mtime:      time.Now(),
//			},
//		}
//	case FILL_MATRIX_SELECTOTR:
//	}
//
//
//	return multiOptions
//}