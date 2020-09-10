package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"log"
	"officialsummary/models"
)

type CategoryAddController struct {
	beego.Controller
}
func (this *CategoryAddController) CategoryAddDis() {
	pathurl := this.Ctx.Input.Param(":id")
	fmt.Println(pathurl)
	release_version := this.GetString("Releaseversion")
	log.Printf(release_version)
	table,err:= new(models.JobInfoList).QueryJob(pathurl)
	if err!=nil{
		logs.Error("SummaryVersionController => ", err)
		this.Abort("404")
	}
	fmt.Println(table)
	this.Data["version"] = pathurl
	this.Data["Contents"]=table
	this.TplName = "category/categoryadd.html"

}

func (this *CategoryAddController) CategoryAdd() {
	//pathurl := this.Ctx.Input.Param(":id")
	//fmt.Println(pathurl)
	cate := models.NewJobInfoList()
	catename := this.GetString("CateName")
	log.Println(catename)
	cate.ProjectName = catename
	release_version := this.GetString("Releaseversion")
	log.Printf(release_version)
	cate.ReleaseVersion = release_version
	//fmt.Printf("last update time: %v\n",job.FinishedTime)
    jobname := this.GetStrings("checkBox")
    fmt.Println(jobname)
	//this.Data["version"] = pathurl
	if err := cate.Insert(); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Insert data fail"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Insert data success"}
	}
	fmt.Printf("after update, the job is: %v\n", cate)
	//this.ServeJSON()
	this.Redirect("/category/"+release_version+"", 302)


}