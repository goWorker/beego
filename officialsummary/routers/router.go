package routers

import (
	"officialsummary/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/summary/:all",&controllers.SummaryVersionController{})
	//beego.Router("/", &controllers.HomeController{}, "get:Index")
	//beego.Router("/index", &controllers.HomeController{}, "get:Index")

	beego.Router("/jenkins",&controllers.HomeController{},"post:ReceiveDataFromJenkins")
}
