package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"officialsummary/models"
)

type JobHistoryController struct {
	beego.Controller
}

func (this *JobHistoryController) GetJobHistory() {

	pathurl := this.Ctx.Input.Param(":all")
	var s string = ""
	pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(pathurl)
	ID := this.Ctx.Input.Param(":id")
	table,err:= new(models.JobList).SearchJobHistory(pathurl,ID)
	if err!=nil{
		logs.Error("JobModifyController => ", err)
		this.Abort("404")
	}
	//ver := (beego.AppConfig.String("versionList"))
	//version := strings.Split(ver,"||")
	//fmt.Println(version)
	this.Data["Contents"]=table
	this.Data["version"] = pathurl
	//this.Data["versionList"] = version
	this.TplName = "jobhistory.html"

}