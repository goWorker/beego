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

	table, err := new(models.JobList).HomeData()
	if err != nil {
		logs.Error("HomeController.Index => ", err)
		c.Abort("404")
	}
	c.Data["Contents"] = table
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
	url := c.GetString("url")
	fmt.Println(url)
	err, detail := utils.SwitchXMLToStruct(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.ResponseWriter.Status = 200
	c.Ctx.WriteString("Parse URL OK!\n")

	jl := models.NewJobList()
	re, _ := regexp.Compile("http://172.25.153.50:8080/view/(.*)/job/(.*)/(.*)/robot/report/output.xml")

	submatch := re.FindSubmatch([]byte(url))
	releasNO := string(submatch[1])
	jobName := string(submatch[2])
	buildNo := string(submatch[3])
	jl.ReleaseVersion = releasNO
	jl.JobName = jobName
	jl.LogUrl = "http://172.25.153.50:8080/view/" + releasNO + "/job/" + jobName + "/" + buildNo + "/robot/report/log.html"

	//fmt.Println(detail.Generated)
	//here should write the location to conf
	location, err := time.LoadLocation("America/New_York")
	timeprune, err := time.ParseInLocation("20060102 15:04:05", detail.Generated, location)
	fmt.Printf("the local time is : %v\n", timeprune)
	if err != nil {
		fmt.Println("Get finished time error")
	}
	jl.FinishedTime = timeprune.In(time.UTC)
	fmt.Printf("The UTC Time is: %v", jl.FinishedTime)

	for i := 0; i < len(detail.Suite); i++ {
		fmt.Println(detail.Suite[i].Status)
		jl.Status = detail.Suite[i].Status.Status
		//for j := 0; j<len(detail.Suite[i].Metadata.ItemList);j++{
		//	jl.Build = detail.Suite[i].Metadata.ItemList[j].Build_Owner_Tag
		//}
		//fmt.Println("The item is:\n")
		if len(detail.Suite[i].Metadata.Item) > 0 {
			if len(detail.Suite[i].Metadata.Item[0].Item) > 0 {
				jl.Build = detail.Suite[i].Metadata.Item[0].Item
			} else {
				jl.Build = "empty"

			}
			if len(detail.Suite[i].Metadata.Item[1].Item) > 0 {
				jl.Owner = detail.Suite[i].Metadata.Item[1].Item
			} else {
				jl.Owner = "empty"

			}
			if len(detail.Suite[i].Metadata.Item[2].Item) > 0 {
				jl.Tag = detail.Suite[i].Metadata.Item[2].Item
			} else {
				jl.Tag = "empty"

			}
		} else {
			jl.Build = "empty"
			jl.Owner = "empty"
			jl.Tag = "empty"
		}

		//jl.Build = detail.Suite[i].Metadata.ItemList[0].Build_Owner_Tag
		//jl.Owner = detail.Suite[i].Metadata.ItemList[1].Build_Owner_Tag
		//jl.Tag = detail.Suite[i].Metadata.ItemList[2].Build_Owner_Tag

		//switch detail.Suite[i].Status.Status{
		//case "PASS":
		//	jl.Status=1
		//case "FAIL":
		//	jl.Status = 0
		//default:
		//	jl.Status = 2
		//}

	}
	for h := 0; h < len(detail.Statistics); h++ {
		//fmt.Println(robot.Statistics[h].Totals)
		for st := 0; st < len(detail.Statistics[h].Totals); st++ {
			//fmt.Println(robot.Statistics[h].Totals[st].States)
			for m := 0; m < len(detail.Statistics[h].Totals[st].States); m++ {
				fmt.Println(detail.Statistics[h].Totals[st].States[m].Stat)
				if detail.Statistics[h].Totals[st].States[m].Stat == "All Tests" {
					fmt.Println(detail.Statistics[h].Totals[st].States[m].Pass)
					fmt.Println(detail.Statistics[h].Totals[st].States[m].Fail)
					jl.PassNum, err = strconv.Atoi(detail.Statistics[h].Totals[st].States[m].Pass)
					jl.FailNum, err = strconv.Atoi(detail.Statistics[h].Totals[st].States[m].Fail)
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
