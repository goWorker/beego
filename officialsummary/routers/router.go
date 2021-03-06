package routers

import (
	"officialsummary/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/summary/add",&controllers.JobAddController{},"get:JobAddDis")
	beego.Router("/summary/add",&controllers.JobAddController{},"post:JobAdd")
	beego.Router("/summary/:all",&controllers.SummaryVersionController{})
	beego.Router("/category/:id/add",&controllers.CategoryAddController{},"get:CategoryAddDis")
	beego.Router("/category/:id/add",&controllers.CategoryAddController{},"post:CategoryAdd")
	beego.Router("/category/:id/:all",&controllers.CategoryJobDisplayController{},"get:CategoryJobDis")
	beego.Router("/category/:all",&controllers.CategorySummController{})
	beego.Router("/category/:id/history/",&controllers.CategoryHistoryController{}, "get:GetDir")
	beego.Router("/category/:id/aggregate/",&controllers.CategoryAggController{}, "post:CategoryAggragate")
	beego.Router("/category/:all/delete/:id",&controllers.CateDeleteController{},"post:Delete")
	beego.Router("/category/:all/edit/:id",&controllers.CateEditController{},"get:GetCateDetail")
	beego.Router("/category/:all/edit/:id",&controllers.CateEditController{},"post:CategoryEdit")
	beego.Router("/summary/:all/history/:id", &controllers.JobHistoryController{},"get:GetJobHistory")
	beego.Router("/summary/:all/history/edit/:id", &controllers.JobHistoryModifyController{},"get:GetJobDetail")
	beego.Router("/summary/:all/history/edit/:id", &controllers.JobHistoryModifyController{},"post:ModifyHisJob")
	beego.Router("/summary/:all/history/delete/:id",&controllers.JobHistoryModifyController{},"post:HisJobDelete")
	beego.Router("/summary/:all/edit/:id",&controllers.JobModifyController{},"get:GetJobDetail")
	beego.Router("/summary/:all/edit/:id",&controllers.JobModifyController{},"post:ModifyJob")
	beego.Router("/summary/:all/delete/:id",&controllers.JobDeleteController{},"post:Delete")
	beego.Router("/jenkins",&controllers.HomeController{},"post:ReceiveDataFromJenkins")
}
