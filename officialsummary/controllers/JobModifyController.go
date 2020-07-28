package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"log"
	"officialsummary/models"
	"strings"
	//"time"
)

type JobModifyController struct {
	beego.Controller
}

func (this *JobModifyController) GetJobDetail() {

	pathurl := this.Ctx.Input.Param(":all")
	var s string = ""
	pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(pathurl)
	ID := this.Ctx.Input.Param(":id")
	table,err:= new(models.JobList).SearchJob(pathurl,ID)
	if err!=nil{
		logs.Error("JobModifyController => ", err)
		this.Abort("404")
	}
	ver := (beego.AppConfig.String("versionList"))
	version := strings.Split(ver,"||")
	fmt.Println(version)
	this.Data["Contents"]=table
	this.Data["version"] = pathurl
	this.Data["versionList"] = version
	this.TplName = "jobmodify.html"

}

func (this *JobModifyController) ModifyJob() {
	id ,_ := this.GetInt("Id")
	fmt.Println("post id: ",id)
	job,err := models.NewJobList().SelectByDocId(id)
	release_version := this.GetString("Releaseversion")
	job.ReleaseVersion=release_version
	jobname := this.GetString("JobName")
	fmt.Println("this is jobname:",jobname)
	status := this.GetString("Status")
	job.Status=status
	passnum,err := this.GetInt("PassNum")
	if err !=nil{
		log.Println(err)
	}
	job.PassNum=passnum
	failnum,err := this.GetInt("FailNum")
	if err !=nil{
		log.Println(err)
	}
	job.FailNum =failnum
	exenum,err := this.GetInt("ExeNum")
	if err !=nil{
		log.Println(err)
	}
	job.ExeNum = exenum
	debugpending := this.GetString("DebugPending")
	job.DebugPending = debugpending
	//source := this.GetString("Source")
	tag := this.GetString("Tag")
	job.Tag = tag
	comment := this.GetString("Comment")
	job.Comment = comment

	//lastupdate := this.GetString("LastUpdate")
	owner := this.GetString("Owner")
	job.Owner=owner
	logurl := this.GetString("LogUrl")
	job.LogUrl = logurl
	//job, err := models.NewJobList().SelectByDocId(Id)
	//job := models.JobList{id,jobname,release_version,status,passnum,failnum,exenum,debugpending,tag,source,comment,owner,logurl,time.Now()}
	if _, err := job.InsertOrUpdate(); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}
	this.ServeJSON()




}
