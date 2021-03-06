package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"officialsummary/models"
	"strings"
)

type SummaryVersionController struct {
	beego.Controller
}


func (this *SummaryVersionController) Get() {

	pathurl := this.Ctx.Input.Param(":all")
	var s string = ""
	pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(pathurl)
	table,err:= new(models.JobList).SummaryForVersionData(pathurl)
	if err!=nil{
		logs.Error("SummaryVersionController => ", err)
		this.Abort("404")
	}
	ver := (beego.AppConfig.String("versionList"))
	version := strings.Split(ver,"||")
	fmt.Println(version)
	this.Data["Contents"]=table
	this.Data["version"] = pathurl
	this.Data["versionList"] = version
	this.Data["subBar"] = "summary"
	this.TplName = "jobsummary.html"

}

