package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
func GetSelector(selector_id int) (Selector){
	o := orm.NewOrm()
	selector:=Selector{Id:selector_id}
	o.Read(&selector)
	o.QueryTable(Option{}).Filter("Selector__id", selector.Id).All(&selector.Option)

	return selector
}
func AddSelector(title string,question_id , template_type int) (Selector){
	o := orm.NewOrm()

	selector:=Selector{
		Title:title,
		Question:&Question{Id:question_id},
		TemplateType:int8(template_type),
	}

	selector_id,err:=o.Insert(&selector)

	if err == nil {
		options := selectorTemplate(int(selector_id),template_type,2)
		o.InsertMulti(len(options),options)
	}

	selector = GetSelector(selector.Id)
	return selector
}

func AddOption(selector_id , template_type int) (Option){
	options := []Option{}
	o := orm.NewOrm()
	options = selectorTemplate(selector_id,template_type,1)
	option := options[0]
	o.Insert(&option)

	fmt.Println(option)
	return option
}

func selectorTemplate(selector_id int,template_type int,createNum int)([]Option){
	if createNum == 0 {
		createNum = 1
	}
	multiOptions :=make([]Option,0)

	switch template_type{
	case SINGLE_SELECTOTR:
		for i:=0;i<createNum;i++{
			opt := Option{
				Selector: &Selector{Id:selector_id},
				Title:      "选项标题",
			}
			multiOptions = append(multiOptions,opt)
		}
	case MULTI_SELECTOTR:
		for i:=0;i<createNum;i++{
			opt := Option{
				Selector: &Selector{Id:selector_id},
				Title:      "选项标题",
			}
			multiOptions = append(multiOptions,opt)
		}
	case SCORE_SELECTOTR:
		multiOptions= []Option{
			Option{
				Selector: &Selector{Id:selector_id},
				Title:      "选项标题",
			},
		}
	case SCORE_MATRIX_SELECTOTR:
	case FILL_SELECTOTR:
		multiOptions= []Option{
			Option{
				Selector: &Selector{Id:selector_id},
				Title:      "选项标题",
			},
		}
	case FILL_MATRIX_SELECTOTR:
	}


	return multiOptions
}