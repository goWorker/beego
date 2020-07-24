package controllers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"officialsummary/models"
	"officialsummary/utils"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {

	table,err:= new(models.JobList).HomeData()
	if err!=nil{
		logs.Error("HomeController.Index => ", err)
		c.Abort("404")
	}
	c.Data["Contents"]=table
		c.TplName = "jobsummary.html"

}
//func (c *HomeController) SummaryForVersion() {
//
//	path := c.Ctx.Input.Param(":id")
//	var s string = "'"
//	path = fmt.Sprintf("%s%s%s",s,path,s)
//	fmt.Println(path)
//	table,err:= new(models.JobList).SummaryForVersionData(path)
//	if err!=nil{
//		logs.Error("HomeController.SummaryForVersion => ", err)
//		c.Abort("404")
//	}
//	c.Data["Contents"]=table
//	c.TplName = "600.html"
//
//}

func (c *HomeController) ReceiveDataFromJenkins() {
	url:=c.GetString("url")
	fmt.Println(url)
	err,detail := utils.SwitchXMLToStruct(url)
	if err != nil{
		fmt.Println(err)
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Status = 200
	c.Ctx.WriteString("Parse URL OK!\n")

	jl := models.NewJobList()
	re,_:= regexp.Compile("http://172.25.153.50:8080/view/(.*)/job/(.*)/(.*)/robot/report/output.xml")
	submatch := re.FindSubmatch([]byte(url))
	releasNO := string(submatch[1])
	jobName := string(submatch[2])
	jl.ReleaseVersion = releasNO
	jl.JobName = jobName

	fmt.Println(detail.Generated)
	jl.FinishedTime,err = time.Parse("20060102 15:04:05.000",detail.Generated)
	fmt.Println(time.Parse("2006-01-02T15:04:05.000Z",detail.Generated))
	if err != nil{
		fmt.Println("Get finished time error")
	}
	for i := 0; i <len(detail.Suite); i++{
		fmt.Println(detail.Suite[i].Status)
		jl.Status=detail.Suite[i].Status.Status

			//switch detail.Suite[i].Status.Status{
			//case "PASS":
			//	jl.Status=1
			//case "FAIL":
			//	jl.Status = 0
			//default:
			//	jl.Status = 2
			//}



	}
	for h := 0; h < len(detail.Statistics); h++{
		//fmt.Println(robot.Statistics[h].Totals)
		for st := 0; st < len(detail.Statistics[h].Totals);st++{
			//fmt.Println(robot.Statistics[h].Totals[st].States)
			for m :=0;m<len(detail.Statistics[h].Totals[st].States);m++{
				fmt.Println(detail.Statistics[h].Totals[st].States[m].Stat)
				if detail.Statistics[h].Totals[st].States[m].Stat == "All Tests"{
					fmt.Println(detail.Statistics[h].Totals[st].States[m].Pass)
					fmt.Println(detail.Statistics[h].Totals[st].States[m].Fail)
					jl.PassNum,err = strconv.Atoi(detail.Statistics[h].Totals[st].States[m].Pass)
					jl.FailNum,err = strconv.Atoi(detail.Statistics[h].Totals[st].States[m].Fail)
					jl.ExeNum = jl.PassNum + jl.FailNum
				}
			}
		}
	}
	jl.Source = "Jenkins"
	if err := jl.Insert(); err != nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.Status = 400
		c.Ctx.WriteString("Insert DB error")
	}

}


