package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	Status 	[]SuiteStatus		`xml:"status"`

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



func ConvertTimeToSeconds(timestring string) int64{
	tm2, _ := time.Parse("20060102 03:04:05", timestring)

	//fmt.Println(tm2.Unix())
	return tm2.Unix()
}
//func (at *AutoSummary)ConvertXML(urls string)
func ConvertXML(urls string) {
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
	if err != nil {
		return
	}

	var robot Robot
	xml.Unmarshal(byteValue, &robot)
	//fmt.Println(robot)
	for j := 0; j < len(robot.Statistics); j++{
		fmt.Println(robot.Statistics)
	}
	//for i := 0; i < len(robot.Suite); i++ {
	//	for suite := 0; suite < len(robot.Suite[i].Suite); suite++ {
	//		for test := 0; test < len(robot.Suite[i].Suite[suite].Test); test++ {
	//			//fmt.Println(robot.Suite[i].Suite[suite].Test[test].Name)
	//			at.CaseName = robot.Suite[i].Suite[suite].Test[test].Name
	//
	//			//fmt.Println(robot.Suite[i].Suite[suite].Test[test].Tags)
	//			for tag := 0; tag <len(robot.Suite[i].Suite[suite].Test[test].Tags);tag++{
	//				casetagstr := strings.Replace(strings.Trim(fmt.Sprint(robot.Suite[i].Suite[suite].Test[test].Tags[tag].Tag), "[]"), " ", ",", -1)
	//				fmt.Println(casetagstr)
	//				at.CaseTag=robot.Suite[i].Suite[suite].Test[test].Tags[tag].Tag
	//			}
	//			for status := 0; status < len(robot.Suite[i].Suite[suite].Test[test].Status); status++ {
	//				//fmt.Println(robot.Suite[i].Suite[suite].Test[test].Status[status].Status)
	//				at.Status = robot.Suite[i].Suite[suite].Test[test].Status[status].Status
	//				starttime := robot.Suite[i].Suite[suite].Test[test].Status[status].Starttime
	//				endtime := robot.Suite[i].Suite[suite].Test[test].Status[status].Endtime
	//				at.ExecuteTime = ConvertTimeToSeconds(endtime) - ConvertTimeToSeconds(starttime)
	//
	//			}
	//
	//		}
	//
	//	}
	//
	//	fmt.Println(at.CaseTag)
	//	fmt.Println(at.CaseName)
	//	fmt.Println(at.ExecuteTime)
	//}
}
func main(){
	//ConvertXML(url)
	url := "http://172.25.153.50:8080/view/6.0.0/job/Felix_BasicAndSpring/60/robot/report/output.xml"
	re,_:= regexp.Compile("http://172.25.153.50:8080/view/(.*)/job/(.*)/(.*)/robot/report/output.xml")

	submatch := re.FindSubmatch([]byte(url))
	//fmt.Println("FindSubmatch", submatch)
	//for _, v := range submatch {
	//	fmt.Println(string(v))
	//}
	releasNO := string(submatch[1])
	jobName := string(submatch[2])
	fmt.Println(releasNO)
	fmt.Println(jobName)

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
