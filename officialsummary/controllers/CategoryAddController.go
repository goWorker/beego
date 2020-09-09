package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"officialsummary/models"
)

type CategoryAddController struct {
	beego.Controller
}
func (this *CategoryAddController) CategoryAddDis() {

	this.TplName = "category/categoryadd.html"

}

func (this *CategoryAddController) CategoryAdd() {

	cate := models.NewJobInfoList()
	catename := this.GetString("CateName")
	log.Println(catename)
	cate.ProjectName = catename



	release_version := this.GetString("Releaseversion")
	log.Printf(release_version)
	cate.ReleaseVersion = release_version
	//fmt.Printf("last update time: %v\n",job.FinishedTime)

	if err := cate.Insert(); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Insert data fail"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "Insert data success"}
	}
	fmt.Printf("after update, the job is: %v\n", cate)
	//this.ServeJSON()
	this.Redirect("/category/"+release_version+"", 302)


}