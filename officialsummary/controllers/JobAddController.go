package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"officialsummary/models"
	"time"
)

type JobAddController struct {
	beego.Controller
}
func (this *JobAddController) JobAddDis() {

	//pathurl := this.Ctx.Input.Param(":all")
	//var s string = ""
	//pathurl = fmt.Sprintf("%s%s%s",s,pathurl,s)
	//fmt.Println(pathurl)
	//ID := this.Ctx.Input.Param(":id")
	//table,err:= new(models.JobList).SearchJob(pathurl,ID)
	//if err!=nil{
	//	logs.Error("JobModifyController => ", err)
	//	this.Abort("404")
	//}
	//ver := (beego.AppConfig.String("versionList"))
	//version := strings.Split(ver,"||")
	//fmt.Println(version)
	//this.Data["Contents"]=table
	//this.Data["version"] = pathurl
	//this.Data["versionList"] = version

	this.TplName = "jobadd.html"


}

func (this *JobAddController) JobAdd() {

	//id ,_ := this.GetInt("Id")
	//fmt.Println("post id: ",id)
	job := models.NewJobList()
	jobname := this.GetString("JobName")
	log.Println(jobname)
	job.JobName = jobname

	release_version := this.GetString("Releaseversion")
	log.Printf(release_version)
	job.ReleaseVersion = release_version
	//fmt.Printf("last update time: %v\n",job.FinishedTime)
	status := this.GetString("Status")
	log.Println(status)
	job.Status = status
	passnum, err := this.GetInt("PassNum")
	if err != nil {
		log.Println(err)
		log.Println(passnum)}
	failnum, err := this.GetInt("FailNum")
	if err != nil {
		log.Println(err)
	}
	log.Println(failnum)

	exenum, err := this.GetInt("ExeNum")
	if err != nil {
		log.Println(err)
	}
	log.Println(exenum)
	job.PassNum = passnum
	job.FailNum = failnum
	job.ExeNum = exenum
	debugpending := this.GetString("DebugPending")
	log.Printf("debug pending: %v",debugpending)
	job.DebugPending = debugpending
	//source := this.GetString("Source")
	tag := this.GetString("Tag")
	job.Tag = tag
	comment := this.GetString("Comment")
	job.Comment = comment
	owner := this.GetString("Owner")
	job.Owner = owner
	logurl := this.GetString("LogUrl")
	job.LogUrl = logurl

	lastupdate := this.GetStrings("UTCDateTime")
	fmt.Printf("last update: %s\n",lastupdate[0])
	//finishedTime := time.Now()
	var timeLayoutStr = "2006-01-02 15:04:05"
	st, _ := time.Parse(timeLayoutStr, lastupdate[0]) //string转time
	fmt.Println(st)
	job.FinishedTime = st
	//log.Printf("finish time: %s\n",finishedTime)
	if err := job.Insert(); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Insert data fail"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Insert data success"}
	}
	fmt.Printf("after update, the job is: %v\n", job)
	//this.ServeJSON()
	this.Redirect("/summary/"+release_version+"", 302)


}