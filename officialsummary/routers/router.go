package routers

import (
	"officialsummary/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/summary/:all",&controllers.SummaryVersionController{})
	//beego.Router("/", &controllers.HomeController{}, "get:Index")
	//beego.Router("/index", &controllers.HomeController{}, "get:Index")
	beego.Router("/summary/:all/edit/:id",&controllers.JobModifyController{},"get:GetJobDetail")
	beego.Router("/summary/:all/edit/:id",&controllers.JobModifyController{},"post:ModifyJob")
	beego.Router("/jenkins",&controllers.HomeController{},"post:ReceiveDataFromJenkins")
}
