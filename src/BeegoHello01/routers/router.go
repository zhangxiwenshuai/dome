package routers

import (
	"BeegoHello01/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterContreller{})
    beego.Router("/", &controllers.MainController{})
}
