package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)



type Robot struct{
	//XMLName		xml.Name	`xml:"robot"`
	Suite 	[]Suite		`xml:"suite"`
	Generated 	string 		`xml:"generated,attr"`
	Statistics	[]Statistics	`xml:"statistics"`
}
type Statistics struct {
	XMLName		xml.Name	`xml:"statistics"`
	Totals 	[]Total `xml:"total"`

}
type Total struct{
	States 	[]Stats		`xml:"stat"`
}
type Stats struct{
	//XMLName		xml.Name	`xml:stat`
	Fail 	string 	`xml:"fail,attr"`
	Pass 	string 	`xml:"pass,attr"`
	Stat 	string 	`xml:",innerxml"`
}


type Suite struct{
	//XMLName		xml.Name	`xml:suite`
	Suite	[]NextSuite 	`xml:"suite"`
	Status 	SuiteStatus		`xml:"status"`

}

type SuiteStatus struct{
	Status	string	`xml:"status,attr"`
}

type NextSuite struct{
	//XMLName		xml.Name	`xml:suite`
	Test 	[]Test	`xml:"test"`
	//Name	string 	`xml:"name,attr"`
}
type Test struct {
	//XMLName		xml.Name 	`xml:"test"`
	Name 	string 	`xml:"name,attr"`
	Status 	[]Status 	`xml:"status"`
	Tags 	[]Tags 		`xml:"tags"`
}
type Tags struct {
	Tag 	[]string 	`xml:"tag"`
}
type Status struct {
	//XMLName		xml.Name 	`xml:"status"`
	Status 	string 		`xml:"status,attr"`
	Endtime 	string	`xml:"endtime,attr"`
	Starttime 	string	`xml:"starttime,attr"`
}
type Res struct{
	Auto 	[]AutoSummary
}
type AutoSummary struct {

	CaseName       string    `orm:"size(100);null" json:"case_name"`
	CaseTag       []string    `orm:"size(200);null" json:"case_tag"`
	Build    string    `orm:"size(100);null" json:"build"`
	ExecuteTime          int64    `orm:"size(100);null" json:"execute_time"`
	ExecuteDate		string	`orm:"auto_now;type(date)" json:"execute_date"`
	Status         string    `orm:"size(100);null" json:"status"`
	Pass 		string
	Fail 		string


}
type JobList struct{
	Id 				int 		`orm:"pk;auto;size(11)" json:"id"`
	JobName 		string		`orm:"size(100)" json:"jobName"`
	ReleaseVersion	string 		`orm:"size(100)" json:"releaseVersion"`
	Status			int			`orm:"size(10);default(2)" json:"status"`
	PassNum			int			`orm:"size(10)" json:"passNum"`
	FailNum			int			`orm:"size(10)" json:"failNum"`
	ExeNum			int			`orm:"size(10)" json:"exeNum"`
	DebugPending	string 		`orm:"size(100);null" json:"debugPending"`
	Tag 			string		`orm:"size(100);null" json:"tag"`
	Source 			string		`orm:"size(100);null" json:"source"`
	Comment 		string		`orm:"size(100);null" json:"comment"`
	Owner			string		`orm:"size(100);null" json:"owner"`
	LogUrl			string		`orm:"size(100);null" json:"log_url"`
	FinishedTime 	time.Time	`orm:"default(auto_now_add);type(datetime)" json:"finished_time"`
}


func ConvertTimeToSeconds(timestring string) int64{
	tm2, _ := time.Parse("20060102 03:04:05", timestring)

	//fmt.Println(tm2.Unix())
	return tm2.Unix()
}
//func (at *AutoSummary)ConvertXML(urls string)
func ConvertXML(urls string) (){
	resp, err := http.Get(urls)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}
	byteValue, _ := ioutil.ReadAll(resp.Body)
	var robot Robot
	err=xml.Unmarshal(byteValue,&robot)
	if err != nil {
		return
	}
	jl := JobList{}
	re,_:= regexp.Compile("http://172.25.153.50:8080/view/(.*)/job/(.*)/(.*)/robot/report/output.xml")
	submatch := re.FindSubmatch([]byte(urls))
	releasNO := string(submatch[1])
	jobName := string(submatch[2])
	jl.ReleaseVersion = releasNO
	jl.JobName = jobName

	fmt.Println(robot.Generated)
	jl.FinishedTime,err = time.Parse("2006-01-02T15:04:05.000Z",robot.Generated)
	for i := 0; i <len(robot.Suite); i++{
		fmt.Println(robot.Suite[i].Status)
		//jl.Status = detail.Suite[i].Status.Status

			switch robot.Suite[i].Status.Status{
			case "PASS":
				jl.Status=1
			case "FAIL":
				jl.Status = 0
			default:
				jl.Status = 2
			}



	}
	for h := 0; h < len(robot.Statistics); h++{
		//fmt.Println(robot.Statistics[h].Totals)
		for st := 0; st < len(robot.Statistics[h].Totals);st++{
			//fmt.Println(robot.Statistics[h].Totals[st].States)
			for m :=0;m<len(robot.Statistics[h].Totals[st].States);m++{
				fmt.Println(robot.Statistics[h].Totals[st].States[m].Stat)
				if robot.Statistics[h].Totals[st].States[m].Stat == "All Tests"{
					fmt.Println(robot.Statistics[h].Totals[st].States[m].Pass)
					fmt.Println(robot.Statistics[h].Totals[st].States[m].Fail)
					jl.PassNum,err = strconv.Atoi(robot.Statistics[h].Totals[st].States[m].Pass)
					jl.FailNum,err = strconv.Atoi(robot.Statistics[h].Totals[st].States[m].Fail)
				}
			}
		}
	}
	fmt.Println(jl)
}
func main(){
	//ConvertXML(url)
	url := "http://172.25.153.50:8080/view/6.0.0/job/Felix_BasicAndSpring/60/robot/report/output.xml"
	//re,_:= regexp.Compile("http://172.25.153.50:8080/view/(.*)/job/(.*)/(.*)/robot/report/output.xml")
	//
	//submatch := re.FindSubmatch([]byte(url))
	ConvertXML(url)
	//releasNO := string(submatch[1])
	//jobName := string(submatch[2])
	//fmt.Println(releasNO)
	//fmt.Println(jobName)
	//at := AutoSummary{}
	//at.ConvertXML(url)

	////定义和上面的FindIndex一样
	//submatchindex := re.FindSubmatchIndex([]byte(url))
	//fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配

	//FindAllSubmatchIndex,查找所有字匹配的index

}
//func main(){
//	at := AutoSummary{}
//	at.ConvertXML(url)
//}
