package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Robot struct{
	//XMLName		xml.Name	`xml:"robot"`
	Suite 	[]Suite		`xml:"suite"`
	Generated 	string 		`xml:"generated,attr"`
	//Statistics	[]Statistics	`xml:"statistics"`
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
	Test 	[]Test	`xml:"test""`
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

func SwitchXMLToStruct(url string)(err error,result Robot) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	byteValue, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var robot Robot
	err=xml.Unmarshal(byteValue,&robot)
	if err != nil{
		return
	}
	fmt.Println(robot)
	return err,robot
}