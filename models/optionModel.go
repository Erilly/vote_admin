package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func AddOption(selector_id , template_type int) (Option){
	options := []Option{}
	o := orm.NewOrm()
	options = selectorTemplate(selector_id,template_type,1)
	option := options[0]
	o.Insert(&option)

	fmt.Println(option)
	return option
}
func UpdateOption(updateData *Option) (int64){
	o := orm.NewOrm()
	fmt.Print(updateData)
	num,_ :=o.Update(updateData)
	return num
}
func DeleteOption(where *Option) (int64){
	o := orm.NewOrm()
	num,_ := o.Delete(where)
	return num
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