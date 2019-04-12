package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"vote_admin/models"
	_ "vote_admin/routers"
)

func init(){
	//注册数据库
	models.RegisterDB()
}

func main() {
	//是否开启ORM调试
	orm.Debug = false
	//自动建表
	orm.RunSyncdb("default", false, true)

	//模板中使用{{add $index}}或{{$index|add}}
	beego.AddFuncMap("add", Indexaddone)

	//模板中使用{{percent $A $B}}
	beego.AddFuncMap("percent", Percent)

	//模板中使用{{anwerCount question_d}}
	beego.AddFuncMap("anwercount", AnwerCount)

	beego.Run()
}

//自定义模板函数，索引加1
func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}

//自定义模板函数，求百分比
func Percent(a,b int) (string) {
	if b>0 && a>0 {
		af:=float32(a)
		bf:=float32(b)
		percent:=(af/bf)*100
		value := fmt.Sprintf("%.2f", percent)
		return value
	}
	return "0"
}

//获取提交答卷数量
func AnwerCount(question_id int) int {
	return models.AnserLogCount(question_id)
}
