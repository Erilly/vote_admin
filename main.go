package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "vote_admin/routers"
	"github.com/astaxie/beego/orm"
	"vote_admin/models"
	"github.com/astaxie/beego"
)

func init(){
	//注册数据库
	models.RegisterDB()
}

func main() {
	//是否开启ORM调试
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	//模板中使用{{add $index}}或{{$index|add}}
	beego.AddFuncMap("add", Indexaddone)

	beego.Run()
}

//自定义模板函数，索引加1
func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}
