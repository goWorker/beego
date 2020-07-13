package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"officialsummary/models"
)

type SummaryVersionController struct {
	beego.Controller
}


func (this *SummaryVersionController) Get() {

	pathurl := this.Ctx.Input.Param(":all")
	var s string = "'"
	pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(pathurl)
	table,err:= new(models.JobList).SummaryForVersionData(pathurl)
	if err!=nil{
		logs.Error("SummaryVersionController => ", err)
		this.Abort("404")
	}
	this.Data["Contents"]=table
	this.TplName = "600.html"

}

