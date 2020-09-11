package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"

	//"log"
	"officialsummary/models"
)

type JobDeleteController struct {
	beego.Controller
}

func (this *JobDeleteController) Delete(){
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
	err = job.Delete(version,job.JobName)
	if err != nil{
		return
	}
	//this.TplName = "jobdelete.html"
	this.Redirect("/summary/"+job.ReleaseVersion+"",302)
}
