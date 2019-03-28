package routers

import (
	"github.com/astaxie/beego"
	"vote_admin/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/questionpage-:id([0-9]+)", &controllers.IndexController{})
    beego.Router("/list", &controllers.ListController{})
    beego.Router("/selector", &controllers.SelectorController{})

	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.ListController{})
	beego.AutoRouter(&controllers.SelectorController{})
}
