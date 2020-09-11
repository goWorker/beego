package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"officialsummary/models"

)
type CateDeleteController struct {
	beego.Controller
}
func (this *CateDeleteController) Delete(){
	version := this.Ctx.Input.Param(":all")
	fmt.Printf("version: ",version)
	project_name := this.Ctx.Input.Param(":id")
	fmt.Println(project_name)
	//id, err := strconv.Atoi(artID)
	//fmt.Println("Delete id: ",artID)
	err := new(models.JobInfoList).Delete(version,project_name)

	if err != nil {
		return
	}
	this.Redirect("/category/"+version+"",302)
}