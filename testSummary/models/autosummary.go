package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type AutoSummary struct {
	Id 				int		`orm:"pk;size(11);not null;unique" json:"id"`
	CaseName       string    `orm:"size(100);null" json:"case_name"`
	CaseTag       string    `orm:"size(200);null" json:"case_tag"`
	Build    string    `orm:"size(100);null" json:"build"`
	ExecuteTime          int64    `orm:"size(100);null" json:"execute_time"`
	ExecuteDate		time.Time	`orm:"auto_now;type(datetime)" json:"execute_date"`
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

func NewAutoSummary() *AutoSummary {
	return &AutoSummary{}
}
func (m *AutoSummary) HomeData(pageIndex, pageSize int,fields ...string) (autosummarys []AutoSummary,totalCount int, err error) {
	if len(fields) == 0 {
		fields = append(fields, "id", "case_name", "case_tag", "build", "execute_time","execute_date","status","jira","log","comments")
	}
	sqlFmt := "select id,case_name,case_tag,build,execute_time, execute_date,status, jira,log,comments from " + TNAutoSummary()
	sqlCount := "select count(*) cnt from "+TNAutoSummary()
	fmt.Println(sqlFmt)
	fmt.Println(sqlCount)
	o := orm.NewOrm()
	var params []orm.Params
	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		if len(params) > 0 {
			totalCount, _ = strconv.Atoi(params[0]["cnt"].(string))
			fmt.Println(totalCount)
		}
	}

	if totalCount > 0 {
		_, err = o.Raw(sqlFmt + " limit " + strconv.Itoa(pageSize) + " offset "+ strconv.Itoa((pageIndex-1)*pageSize)).QueryRows(&autosummarys)
	}
	return
}

func (m *AutoSummary) Select(field string, value interface{}, cols ...string) (autosummary *AutoSummary, err error) {
	if len(cols) == 0 {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m)
	} else {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m, cols...)
	}
	return m, err
}

func (m *AutoSummary) SelectPage(pageIndex, pageSize int) (autosummarys []*AutoSummary, totalCount int, err error) {
	o := orm.NewOrm()
	sql1 := "select count(*) cnt from md_autosummary;"

	err = o.Raw(sql1).QueryRow(&totalCount)
	if err != nil {
		return
	}
	offset := (pageIndex - 1) * pageSize
	sql2 := "select *  from " + TNAutoSummary() +
		" order by case_name limit " + fmt.Sprintf("%d,%d", offset, pageSize)
	_, err = o.Raw(sql2).QueryRows(&autosummarys)
	if err != nil {
		return
	}
	return
}

func (m *AutoSummary) Insert() (err error) {
	if _, err = orm.NewOrm().Insert(m); err != nil {
		return
	}
	return err
}

func (m *AutoSummary) Update(cols ...string) (err error) {
	as := NewAutoSummary()
	as.Id = m.Id
	o := orm.NewOrm()
	if err = o.Read(as); err != nil {
		return err
	}
	_, err = o.Update(m, cols...)
	return err
}