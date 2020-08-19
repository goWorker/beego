package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"log"
	"officialsummary/models"
	"strconv"
)

type JobHistoryModifyController struct {
	beego.Controller
}

func (this *JobHistoryModifyController) GetJobDetail() {

	pathurl := this.Ctx.Input.Param(":all")
	var s string = ""
	pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	fmt.Println(pathurl)
	ID := this.Ctx.Input.Param(":id")
	table,err:= new(models.JobList).SearchHistoryJobDetail(ID)
	if err!=nil{
		logs.Error("JobModifyController => ", err)
		this.Abort("404")
	}
	this.Data["Contents"]=table
	this.TplName = "historyjobmodify.html"

}

func (this *JobHistoryModifyController) ModifyHisJob() {
	id ,_ := this.GetInt("Id")
	fmt.Println("post id: ",id)
	job,err := models.NewJobList().SelectByDocId(id)
	release_version := this.GetString("Releaseversion")
	fmt.Printf("last update time: %v\n",job.FinishedTime)
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
	fmt.Printf("after update, the job is: %v\n",job)
	//this.ServeJSON()
	this.Redirect("/summary/"+job.ReleaseVersion+"/history/"+job.JobName+"",302)

}

func (this *JobHistoryModifyController) HisJobDelete(){
	fmt.Println("test ***************************")
	version := this.Ctx.Input.Param(":all")
	fmt.Printf("version: ",version)
	artID := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(artID)
	fmt.Println("Delete id: ",artID)
	job,err := models.NewJobList().SelectByDocId(id)

	if err != nil {
		return
	}
	//id := strconv.Itoa(artID)
	err = job.DeleteHisJob(artID)
	if err != nil{
		return
	}
	//this.TplName = "jobdelete.html"
	this.Redirect("/summary/"+job.ReleaseVersion+"/history/"+job.JobName+"",302)
}