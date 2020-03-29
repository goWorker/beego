package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	_"github.com/go-sql-driver/mysql"

)

type AutoSummary struct {
	CaseName       string    `orm:"pk;size(100);null;unique" json:"case_name"`
	CaseTag       string    `orm:"size(200);null" json:"case_tag"`
	Build    string    `orm:"size(100);null" json:"build"`
	ExecuteTime          string    `orm:"size(100);null" json:"execute_time"`
	ExecuteDate		time.Time	`orm:"type(datetime)" json:"execute_date"`
	Log         string    `orm:"size(100);null" json:"log"`
	Status         string    `orm:"size(100);null" json:"status"`
	Jira        string       `orm:"size(100);null" json:"jira"`
	Comments 	string       `orm:"size(500);null" json:"comments"`
	BackupStr	string		 `orm:"size(100);null" json:"backup_str"`
	BackupInt	int			`orm:"size(100);null" json:"backup_int"`

}
func (m *AutoSummary) TableName() string {
	return TNAutoSummary()
}

func (m *AutoSummary) HomeData() (autosummaryss []AutoSummary, err error) {
	qs := orm.NewOrm().QueryTable(TNAutoSummary())
	//if pid > -1 {
	//	qs = qs.Filter("pid", pid)
	//}
	//
	//if status == 0 || status == 1 {
	//	qs = qs.Filter("status", status)
	//}
	//qs.OrderBy()
	_, err = qs.OrderBy("-status","case_name").All(&autosummaryss)
	fmt.Println("*************")
	fmt.Println(autosummaryss)
	return
}
