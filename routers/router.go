package routers

import (
	"jjj/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
}
