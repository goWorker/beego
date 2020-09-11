package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"officialsummary/models"
)

type CateEditController struct {
	beego.Controller
}

func (this *CateEditController) GetCateDetail() {

	version := this.Ctx.Input.Param(":all")
	project_name := this.Ctx.Input.Param(":id")
	unselecttable,err:= new(models.JobInfoList).QueryUnSelectedJob(version)

	if err!=nil{
		logs.Error("CateEditController => ", err)
		this.Abort("404")
	}
	selecttable,err := new(models.JobInfoList).QuerySelectedJob(version,project_name)

	if err!=nil{
		logs.Error("CateEditController => ", err)
		this.Abort("404")
	}
	this.Data["ProjectName"]=project_name
	this.Data["Version"]=version
	this.Data["SelectContents"]=selecttable
	this.Data["UnSelectContents"]=unselecttable
	this.TplName = "category/categoryedit.html"

}
func (this *CateEditController) CategoryEdit() {
	version := this.Ctx.Input.Param(":all")
	project_name := this.Ctx.Input.Param(":id")
	err := new(models.JobInfoList).Delete(version,project_name)

	if err != nil {
		return
	}
	checkedjobname := this.GetStrings("checkBox")
	cates:=make([]models.JobInfoList,len(checkedjobname))
	for i := 0; i < len(checkedjobname); i++ {
		cates = append(cates,models.JobInfoList{ProjectName: project_name,ReleaseVersion: version,JobName: checkedjobname[i]})
	}
	fmt.Println(cates)
	if num ,err := orm.NewOrm().InsertMulti(len(cates),cates);err !=nil{
		fmt.Println(err)
	}else {
		fmt.Printf("Insert %d cate\r\n",num)
	}

	uncheckedjobname := this.GetStrings("uncheckBox")
	uncheckcates:=make([]models.JobInfoList,len(uncheckedjobname))
	for i := 0; i < len(uncheckedjobname); i++ {
		uncheckcates = append(uncheckcates,models.JobInfoList{ProjectName: project_name,ReleaseVersion: version,JobName: uncheckedjobname[i]})
	}
	fmt.Println(uncheckcates)

	if num ,err := orm.NewOrm().InsertMulti(len(uncheckcates),uncheckcates);err !=nil{
		fmt.Println(err)
	}else {
		fmt.Printf("Insert %d cate\r\n",num)
	}
	//fmt.Printf("after update, the job is: %v\n", cates)
	this.Redirect("/category/"+version+"", 302)

}