package routers

import (
	"vote_admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/admin", &controllers.AdminController{})
}
