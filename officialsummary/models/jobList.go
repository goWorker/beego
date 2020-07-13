package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type JobList struct{
	Id 				int 		`orm:"pk;auto;size(11)" json:"id"`
	JobName 		string		`orm:"size(100)" json:"jobName"`
	ReleaseVersion	string 		`orm:"size(100)" json:"releaseVersion"`
	Status			string			`orm:"size(50);default('UNEXECUTED')" json:"status"`
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
func (m *JobList) TableName() string {
	return TNjobList()
}

func NewJobList() *JobList {
	return &JobList{}
}
func (m *JobList) HomeData(fields ...string) (joblist []JobList, err error) {
	if len(fields) == 0 {
		fields = append(fields, "id", "job_name", "release_version", "status", "pass_num","fail_num","exe_num","debug_pending","tag","source","comment","owner","log_url","finished_time")
	}
	//sqlFmt := "select id,job_name,release_version,status,pass_num, fail_num,exe_num,debug_pending,tag,source,comment,owner,log_url,finished_time from " + TNjobList()

	sqlFmt := "SELECT a.* FROM "+ TNjobList()+" a,( SELECT job_name, release_version, MAX( finished_time ) ftime FROM jobList GROUP BY job_name, release_version ) b WHERE a.job_name = b.job_name AND a.release_version = b.release_version AND a.finished_time = b.ftime AND a.release_version = '6.1.0'"
	sqlCount := "select count(*) cnt from "+TNjobList()
	fmt.Println(sqlFmt)
	//fmt.Println(sqlCount)
	o := orm.NewOrm()
	var params []orm.Params
	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_, err = o.Raw(sqlFmt).QueryRows(&joblist)

	}

	return
}

func (m *JobList)SummaryForVersionData(version string) (joblist []JobList, err error) {

	sqlFmt := "SELECT a.* FROM "+ TNjobList()+" a,( SELECT job_name, release_version, MAX( finished_time ) ftime FROM jobList GROUP BY job_name, release_version ) b WHERE a.job_name = b.job_name AND a.release_version = b.release_version AND a.finished_time = b.ftime AND a.release_version = " + version
	sqlCount := "select count(*) cnt from "+TNjobList()
	fmt.Println(sqlFmt)
	//fmt.Println(sqlCount)
	o := orm.NewOrm()
	var params []orm.Params

	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_, err = o.Raw(sqlFmt).QueryRows(&joblist)

	}

	return
}

//func (m *JobList) HomeData(pageIndex, pageSize int,fields ...string) (joblist []JobList,totalCount int, err error) {
//	if len(fields) == 0 {
//		fields = append(fields, "id", "job_name", "release_version", "status", "pass_num","fail_num","exe_num","debug_pending","tag","source","comment","owner","log_url","finished_time")
//	}
//	//sqlFmt := "select id,job_name,release_version,status,pass_num, fail_num,exe_num,debug_pending,tag,source,comment,owner,log_url,finished_time from " + TNjobList()
//	sqlFmt := "SELECT a.* FROM "+ TNjobList()+" a,( SELECT job_name, release_version, MAX( finished_time ) ftime FROM jobList GROUP BY job_name, release_version ) b WHERE a.job_name = b.job_name AND a.release_version = b.release_version AND a.finished_time = b.ftime AND a.release_version = '6.1.0'"
//	sqlCount := "select count(*) cnt from "+TNjobList()
//	fmt.Println(sqlFmt)
//	fmt.Println(sqlCount)
//	o := orm.NewOrm()
//	var params []orm.Params
//	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
//		if len(params) > 0 {
//			totalCount, _ = strconv.Atoi(params[0]["cnt"].(string))
//			fmt.Println(totalCount)
//		}
//	}
//
//	if totalCount > 0 {
//		_, err = o.Raw(sqlFmt + " limit " + strconv.Itoa(pageSize) + " offset "+ strconv.Itoa((pageIndex-1)*pageSize)).QueryRows(&joblist)
//	}
//	return
//}
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
