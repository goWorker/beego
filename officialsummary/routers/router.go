package routers

import (
	"officialsummary/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/summary", &controllers.HomeController{}, "get:Index")
	beego.Router("/jenkins",&controllers.HomeController{},"post:ReceiveDataFromJenkins")
}
