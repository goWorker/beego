package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"strings"
	"officialsummary/common"
	"officialsummary/models"
	"officialsummary/utils"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {

	pageIndex, _ := c.GetInt("p", 1)
	//private, _ := c.GetInt("private", 1)
	table,totalCount,err:= new(models.JobList).HomeData(pageIndex,common.PageSize)
	if err!=nil{
		logs.Error("HomeController.Index => ", err)
		c.Abort("404")
	}
	if totalCount > 0{
		c.Data["paginator"]=utils.NewPaginator(c.Ctx.Request,common.PageSize,totalCount)
	}else{
		c.Data["paginator"]=""
	}
	c.Data["Contents"]=table
		c.TplName = "index.html"

}
func (c *HomeController) ReceiveDataFromJenkins() {
	//url:=c.GetString("url")
	//fmt.Println(url)
	//err,detail := utils.SwitchXMLToStruct(url)
	//if err != nil{
	//	return
	//}
	//c.Ctx.ResponseWriter.WriteHeader(200)
	//c.Data["Message"] = "We are getting the "
	//c.TplName = "index.html"
	////fmt.Println(detail)
	//at := models.NewJobList()
	//for i := 0; i < len(detail.Suite); i++ {
	//	for suite := 0; suite < len(detail.Suite[i].Suite); suite++ {
	//		for test := 0; test < len(detail.Suite[i].Suite[suite].Test); test++ {
	//			//fmt.Println(detail.Suite[i].Suite[suite].Test[test].Name)
	//			at.JobName = detail.Suite[i].Suite[suite].Test[test].Name
	//
	//			//fmt.Println(detail.Suite[i].Suite[suite].Test[test].Tags)
	//			for tag := 0; tag <len(detail.Suite[i].Suite[suite].Test[test].Tags);tag++{
	//				casetagstr := strings.Replace(strings.Trim(fmt.Sprint(detail.Suite[i].Suite[suite].Test[test].Tags[tag].Tag), "[]"), " ", ",", -1)
	//				at.Tag=casetagstr
	//			}
	//			for status := 0; status < len(detail.Suite[i].Suite[suite].Test[test].Status); status++ {
	//				//fmt.Println(detail.Suite[i].Suite[suite].Test[test].Status[status].Status)
	//				at.Status = detail.Suite[i].Suite[suite].Test[test].Status[status].Status
	//				starttime := detail.Suite[i].Suite[suite].Test[test].Status[status].Starttime
	//				endtime := detail.Suite[i].Suite[suite].Test[test].Status[status].Endtime
	//				at.ExecuteTime = utils.ConvertTimeToSeconds(endtime) - utils.ConvertTimeToSeconds(starttime)
	//
	//			}
	//			if err := at.Insert(); err != nil {
	//				fmt.Println("Insert db error")
	//			}
	//			at.BackupInt = 0
	//			at.BackupStr = ""
	//			at.Comments = ""
	//			at.Build = ""
	//			at.Jira = ""
	//			at.Log = ""
	//
	//
	//		}
	//
	//	}
	//
	//}

}
//c.CustomAbort(200,"Data received,we are parsing!")


