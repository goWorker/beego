package main

import (
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"io/ioutil"
	"net/http"
)

//xml的struct结构
type Robot struct {
	XMLName	xml.Name	`xml:"robot"`
	Suite 	Suite	`xml:"suite"`
}
type Suite struct {
	Name	string 	`xml:"name,attr"`
	Suite 	NewSuite	`xml:"suite"`
}

type NewSuite struct{
	Name 	string	`xml:"name,attr"`
	Test 	[]string 	`xml:"test"`
}
//发手机短信
func Send_duanxin(url string) (dx Robot, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return dx, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dx, err
	}
    v := Robot{}
	err = xml.Unmarshal(body, &v)//好方便，就这样就解析了xml，so关键是xml的结构
	if err != nil{
		fmt.Printf("error: %v", err)
		return
	}
	//fmt.Println(v)
	fmt.Println(v.Suite)
	return v, err
}
//var aaa = string("<returnsms> <returnstatus>Success</returnstatus><message>ok</message><remainpoint>150528</remainpoint><taskID>518334</taskID><successCounts>1</successCounts></returnsms>")






func main() {
	//_,err := Send_duanxin("http://172.25.153.50:8080/view/6.0.0/job/Allan_BasicAndSpring/43/robot/report/output.xml")
	//if err != nil{
	//	return
	//}
	resp, err := http.Get("http://172.25.153.50:8080/view/6.0.0/job/Allan_BasicAndSpring/43/robot/report/output.xml")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(body);err != nil{
		panic(err)
	}
	//if err := doc.ReadFromFile("/Users/allanyang/Go/src/mycode/output.xml"); err != nil{
	//	panic(err)
	//}
	root := doc.SelectElement("robot")

	for _,suites := range root.SelectElements("suite"){
		for _,suite := range suites.SelectElements("suite"){
			for _, tests := range suite.SelectElements("test"){
				fmt.Println(tests.SelectAttrValue("name","unknown"))
				//for _, status := range tests.SelectElements("status"){
				//	fmt.Println(status.SelectAttrValue("status","unknown"))
				//	fmt.Println(status.SelectAttrValue("starttime","0"))
				//	fmt.Println(status.SelectAttrValue("endtime","0"))
				//	//
				//
				//}
				for _,tags := range tests.SelectElements("tags"){
					fmt.Println(tags.SelectElement("tag").Text())


				}

			}
		}

	}

	//suites	 := root.SelectElement("suite")
	//for _,suite := range suites.SelectElements("suite"){
	//	fmt.Println(suite.SelectElement("test").Tag)
	//	for _, test := range suite.SelectElements("status"){
	//		fmt.Println(test.SelectElement("status").Tag)
	//	}
	//}

}