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

	beego.Run()
}

