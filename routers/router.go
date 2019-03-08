package routers

import (
	"github.com/astaxie/beego"
	"vote_admin/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.IndexController{})
    beego.Router("/list", &controllers.ListController{})
    //beego.Router("/list/edit/:qustionId:string", &controllers.ListController{})

	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.ListController{})
}
