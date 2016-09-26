package routers

import (
	"github.com/Blind238/assemblage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
