package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"log"
	"officialsummary/models"
	"strings"
	"time"
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
	release_version := this.GetString("Releaseversion")
	jobname := this.GetString("JobName")
	fmt.Println("this is jobname:",jobname)
	status := this.GetString("Status")
	passnum,err := this.GetInt("PassNum")
	if err !=nil{
		log.Println(err)
	}
	failnum,err := this.GetInt("FailNum")
	if err !=nil{
		log.Println(err)
	}
	exenum,err := this.GetInt("ExeNum")
	if err !=nil{
		log.Println(err)
	}
	debugpending := this.GetString("DebugPending")
	source := this.GetString("Source")
	tag := this.GetString("Tag")
	comment := this.GetString("Comment")

	//lastupdate := this.GetString("LastUpdate")
	owner := this.GetString("Owner")
	logurl := this.GetString("LogUrl")
	//job, err := models.NewJobList().SelectByDocId(Id)
	job := models.JobList{id,jobname,release_version,status,passnum,failnum,exenum,debugpending,tag,source,comment,owner,logurl,time.Now()}
	if _, err := job.InsertOrUpdate(); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}
	this.ServeJSON()
	//	this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	//} else {
	//	documentStore.DocumentId = int(docId)
	//	if err := documentStore.InsertOrUpdate("markdown", "content"); err != nil {
	//		beego.Error(err)
	//	}
	//}



}
