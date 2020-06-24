package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

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
	LogURL			string		`orm:"size(100);null" json:"logUrl"`
	FinishedTime 	time.Time	`orm:"auto_now;type(datetime);null" json:"finishedTime"`
}
func (m *JobList) TableName() string {
	return TNjobList()
}


func NewJobList() *JobList {
	return &JobList{}
}
func (m *JobList) HomeData(pageIndex, pageSize int,fields ...string) (autosummarys []JobList,totalCount int, err error) {
	if len(fields) == 0 {
		fields = append(fields, "id", "case_name", "case_tag", "build", "execute_time","execute_date","status","jira","log","comments")
	}
	sqlFmt := "select id,case_name,case_tag,build,execute_time, execute_date,status, jira,log,comments from " + TNjobList()
	sqlCount := "select count(*) cnt from "+TNjobList()
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

func (m *JobList) Select(field string, value interface{}, cols ...string) (joblist *JobList, err error) {
	if len(cols) == 0 {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m)
	} else {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m, cols...)
	}
	return m, err
}

func (m *JobList) SelectPage(pageIndex, pageSize int) (autosummarys []*JobList, totalCount int, err error) {
	o := orm.NewOrm()
	sql1 := "select count(*) cnt from md_autosummary;"

	err = o.Raw(sql1).QueryRow(&totalCount)
	if err != nil {
		return
	}
	offset := (pageIndex - 1) * pageSize
	sql2 := "select *  from " + TNjobList() +
		" order by case_name limit " + fmt.Sprintf("%d,%d", offset, pageSize)
	_, err = o.Raw(sql2).QueryRows(&autosummarys)
	if err != nil {
		return
	}
	return
}

func (m *JobList) Insert() (err error) {
	if _, err = orm.NewOrm().Insert(m); err != nil {
		return
	}
	return err
}

func (m *JobList) Update(cols ...string) (err error) {
	as := NewJobList()
	as.Id = m.Id
	o := orm.NewOrm()
	if err = o.Read(as); err != nil {
		return err
	}
	_, err = o.Update(m, cols...)
	return err
}
