package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"officialsummary/models"
)

type CategoryHistoryController struct {
	beego.Controller
}

func (this *CategoryHistoryController) GetDir() {

	version := this.Ctx.Input.Param(":id")
	//var s string = ""
	//pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(version)
	table,err:= new(models.JobAggregateList).CheckSaveTime(version)
	if err!=nil{
		logs.Error("CategoryHistoryController => ", err)
		this.Abort("404")
	}
	fmt.Println(table)
	this.Data["Contents"]=table
	this.Data["version"] = version
	//this.Data["versionList"] = version
	//this.Data["subBar"] = "cate"
	this.TplName = "category/categoryhistory.html"

}
