package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"officialsummary/models"
)

type CategoryJobDisplayController struct {
	beego.Controller
}


func (this *CategoryJobDisplayController) CategoryJobDis() {

	version := this.Ctx.Input.Param(":id")
	project_name := this.Ctx.Input.Param(":all")
	//var s string = ""
	//pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)

	table,err:= new(models.JobInfoList).CateJobDisplay(version,project_name)
	if err!=nil{
		logs.Error("CategoryJobDisplayController => ", err)
		this.Abort("404")
	}
	fmt.Println(table)
	this.Data["Contents"]=table
	this.TplName = "category/categoryjobdis.html"

}
