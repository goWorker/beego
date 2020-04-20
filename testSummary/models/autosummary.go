package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AutoSummary struct {
	Id 				int		`orm:"pk;size(11);not null;unique" json:"id"`
	CaseName       string    `orm:"size(100);null" json:"case_name"`
	CaseTag       string    `orm:"size(200);null" json:"case_tag"`
	Build    string    `orm:"size(100);null" json:"build"`
	ExecuteTime          string    `orm:"size(100);null" json:"execute_time"`
	ExecuteDate		string	`orm:"auto_now;type(date)" json:"execute_date"`
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
	fmt.Println("&&&&&&&&")

	_, err = qs.OrderBy("execute_date","-status","case_name").All(&autosummaryss)
	fmt.Println("*************")
	fmt.Println(autosummaryss)
	return
}
