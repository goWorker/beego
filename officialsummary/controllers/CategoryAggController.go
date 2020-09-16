package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"officialsummary/models"
)

type CategoryAggController struct {
	beego.Controller
}

func (this *CategoryAggController) CategoryAggragate(){
	version := this.Ctx.Input.Param(":id")
	//var s string = ""
	//pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(version)
	table,err:= new(models.JobInfoList).CateForVersion(version)
	if err!=nil{
		logs.Error("CategoryAggController => ", err)
		this.Abort("404")
	}
	fmt.Println(table)
	jobcatesummlist :=new(models.JobAggregateList).SaveAggregateList(table,version)
	fmt.Println(jobcatesummlist)

	if num ,err := orm.NewOrm().InsertMulti(len(jobcatesummlist),jobcatesummlist);err !=nil{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Save into db error"}
		fmt.Println(err)
	}else {
		fmt.Printf("Insert %d history cate\r\n",num)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Save into db succeed"}
	}
	this.ServeJSON()

}