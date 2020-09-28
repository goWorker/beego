package utils

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Robot struct{
	//XMLName		xml.Name	`xml:"robot"`
	Suite 	[]Suite		`xml:"suite"`
	Generated 	string 		`xml:"generated,attr"`
	Statistics	[]Statistics	`xml:"statistics"`
}
type Statistics struct {
	//XMLName		xml.Name	`xml:"statistics"`
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
	//Suite	[]NextSuite 	`xml:"suite"`
	Status 	SuiteStatus		`xml:"status"`
	Metadata 	ItemList 		`xml:"metadata"`
}
type ItemList struct{
	Item 	[]Item		`xml:"item"`
	//Name 	string		`xml:"name,attr"`
}
type Item struct {
	Item		string 		`xml:",innerxml"`
	//Build_Version	string 	`xml:"Build_Version，attr"`
	//Job_Owner	string 	`xml:",innerxml"`
	//Tag  	string 	`xml:",innerxml"`
}
type SuiteStatus struct{
	Status	string	`xml:"status,attr"`
}

//type NextSuite struct{
//	//XMLName		xml.Name	`xml:suite`
//	Test 	[]Test	`xml:"test""`
//	//Name	string 	`xml:"name,attr"`
//}
//type Test struct {
//	//XMLName		xml.Name 	`xml:"test"`
//	Name 	string 	`xml:"name,attr"`
//	Status 	[]Status 	`xml:"status"`
//	Tags 	[]Tags 		`xml:"tags"`
//}
//type Tags struct {
//	Tag 	[]string 	`xml:"tag"`
//}
//type Status struct {
//	//XMLName		xml.Name 	`xml:"status"`
//	Status 	string 		`xml:"status,attr"`
//	Endtime 	string	`xml:"endtime,attr"`
//	Starttime 	string	`xml:"starttime,attr"`
//}

func SwitchXMLToStruct(url string)(err error,result Robot) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Got err: %v",err)
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
		fmt.Println(err)
		return
	}
	fmt.Println(robot)
	return err,robot

}
var db *sql.DB

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	fmt.Println(sql)
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println("first")
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("second")
		log.Println(err)
		return 0, err
	}
	return count, nil
}