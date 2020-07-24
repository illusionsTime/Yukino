package routers

import (
	"Yukino/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.ErrorController(&controllers.ErrorController{})
}
