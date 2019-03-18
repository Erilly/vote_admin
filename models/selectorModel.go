package models

import (
	"github.com/astaxie/beego/orm"
)

func GetSelector(selector_id int) (Selector,error){
	o := orm.NewOrm()
	selector:=Selector{Id:selector_id}
	err:=o.Read(&selector)

	return selector,err
}

func GetSelectorInfo(selector_id int) (Selector){
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

	selector = GetSelectorInfo(selector.Id)
	return selector
}

func UpdateSelector(updateData *Selector) (int64){
	o := orm.NewOrm()
	num,_ :=o.Update(updateData)
	return num
}

func DeleteSelector(where *Selector) (int64){
	o := orm.NewOrm()
	num,_:=o.Delete(where)

	return num
}
