package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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
	cate := models.JobInfoList{}
	catename := this.GetString("CateName")
	log.Println(catename)
	cate.ProjectName = catename
	release_version := this.GetString("Releaseversion")
	log.Printf(release_version)
	cate.ReleaseVersion = release_version
    jobname := this.GetStrings("checkBox")
    cates:=make([]models.JobInfoList,len(jobname))
	for i := 0; i < len(jobname); i++ {
		cates = append(cates,models.JobInfoList{ProjectName: catename,ReleaseVersion: release_version,JobName: jobname[i]})
	}
	fmt.Println(cates)
	if num ,err := orm.NewOrm().InsertMulti(len(cates),cates);err !=nil{
		fmt.Println(err)
	}else {
		fmt.Printf("Insert %d cate\r\n",num)
	}
	fmt.Printf("after update, the job is: %v\n", cates)
	this.Redirect("/category/"+release_version+"", 302)

}