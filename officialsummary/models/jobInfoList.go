package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type JobInfoList struct {
	Id 				int 		`orm:"pk;auto;size(11)" json:"id"`
	ProjectName		string		`orm:"size(64)" json:"projectName"`
	ReleaseVersion 	string		`orm:"size(32)" json:"releaseVersion"`
	JobName			string		`orm:"size(64)" json:"jobName"`
}
type JobCateSummList struct {
	ProjectName    string		`json:"project_name"`
	TotalJob  	int				`json:"total_job"`
	Executed       int			`json:"executed"`
	Pass           int			`json:"pass"`
	Fail           int			`json:"fail"`
	ExeRatio       float32		`json:"exe_ratio"`
	PassRatio     float32		`json:"pass_ratio"`
	FailRatio     float32		`json:"fail_ratio"`
}
type JobNameList struct {
	JobName		string		`json:"job_name"`
}


func (m *JobInfoList)TableName() string {
	return TNjobInfoList()
}
func NewJobInfoList() *JobInfoList {
	return &JobInfoList{}
}
var updateDely = beego.AppConfig.String("updateDelay")
//fmt.Printf("The update delay is %s days\n",updateDely)

func (m *JobInfoList)CateForVersion(version string) (jobcatesummlist []JobCateSummList, err error) {

	sqlFmt := "SELECT t6.project_name as project_name,\n  t6.totalJob as total_job,\n  ifnull(t5.totalExe, 0) as executed,\n  ifnull(t5.totalPass, 0) as pass,\n  ifnull(t5.totalFail, 0) as fail,\n  ifnull(t5.totalExe, 0)/t6.totalJob as exe_ratio,\n  ifnull(t5.totalPass, 0)/t6.totalJob as pass_ratio,\n  ifnull(t5.totalFail, 0)/t6.totalJob as fail_ratio\n    FROM (SELECT project_name,sum(exe_num) AS totalJob FROM "+TNjobInfoList()+" a,\n                                                            (SELECT t1.job_name,t1.exe_num\n                                                            FROM "+TNjobList()+" t1,\n                                                                 (SELECT job_name, max(id) AS maxId FROM "+TNjobList()+" WHERE release_version = '"+version+"' GROUP BY job_name) t2\n                                                            WHERE t1.id = t2.maxId) b\n    WHERE a.job_name = b.job_name GROUP BY project_name) t6\n        LEFT JOIN\n        (SELECT t4.project_name,sum(t3.exe_num) AS totalExe, sum(t3.pass_num) AS totalPass, sum(t3.fail_num) AS totalFail\n        FROM "+TNjobInfoList()+" t4,\n             (SELECT t1.job_name,t1.exe_num,t1.pass_num,t1.fail_num FROM\n                                                                      "+TNjobList()+" t1,\n                                                                         (SELECT job_name, max(id) AS maxId FROM "+TNjobList()+" WHERE release_version = '"+version+"' AND DATE_FORMAT(finished_time, '%Y-%m-%d') >= DATE_FORMAT(DATE_SUB(NOW(), INTERVAL "+updateDely+" DAY), '%Y-%m-%d') GROUP BY job_name) t2\n             WHERE t1.id = t2.maxId) t3\n        WHERE t4.job_name = t3.job_name  AND t4.release_version = '"+version+"' GROUP BY t4.project_name ) t5\n            ON t6.project_name = t5.project_name;"
	sqlCount := "select count(*) cnt from "+TNjobInfoList()
	o := orm.NewOrm()
	var params []orm.Params

	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_,err = o.Raw(sqlFmt).QueryRows(&jobcatesummlist)

	}
	fmt.Printf("The jobcatesummlist is: %v",jobcatesummlist)

	return
}

func (m *JobInfoList) Insert() (err error) {
	if _, err = orm.NewOrm().Insert(m); err != nil {
		return
	}
	return err
}

func (m *JobInfoList)QueryJob(version string)(jobnamelist []JobNameList,err error){
	//sqlfmt := "select job_name from  "+TNjobList()+" where job_name not in(select job_name from "+TNjobInfoList()+");"
	sqlfmt := "select distinct job_name from  "+TNjobList()+" where release_version='"+version+"' and job_name not in(select job_name from "+TNjobInfoList()+" where release_version='"+version+"');"
	o := orm.NewOrm()
	_,err = o.Raw(sqlfmt).QueryRows(&jobnamelist)
	fmt.Printf("The jobnamelist is: %v\n",jobnamelist)
	return
}
func (jobinfolist *JobInfoList)QueryUnSelectedJob(version string)(jobnamelist []JobNameList,err error){
	sqlfmt := "select distinct job_name from  "+TNjobList()+" where release_version='"+version+"' and job_name not in(select job_name from "+TNjobInfoList()+" where release_version='"+version+"');"
	o := orm.NewOrm()
	_,err = o.Raw(sqlfmt).QueryRows(&jobnamelist)
	fmt.Printf("The jobnamelist is:%v\n",jobnamelist)
	return
}

func (jobinfolist *JobInfoList)QuerySelectedJob(version,project_name string)(jobnamelist []JobNameList,err error){
	sqlfmt := "select job_name from "+TNjobInfoList()+" where release_version='"+version+"' and project_name='"+project_name+"'"
	o := orm.NewOrm()
	_,err = o.Raw(sqlfmt).QueryRows(&jobnamelist)
	fmt.Printf("The jobnamelist is:%v\n",jobnamelist)
	return
}



func (jobinfolist *JobInfoList)Delete(version,project_name string) error{
	o := orm.NewOrm()
	var jobs []*JobInfoList

	_, err := o.QueryTable(jobinfolist.TableName()).Filter("project_name",project_name).Filter("release_version",version ).All(&jobs)
	if err != nil {
		return err
	}
	project_nameid := make([]int, len(jobs))
	for _,item := range jobs{
		project_nameid = append(project_nameid, item.Id)
	}
	_,err =o.QueryTable(new(JobInfoList)).Filter("id__in",project_nameid).Delete()
	if err != nil{
		fmt.Printf("delete User by Ids fail: [%v]\n", err)
	}
	fmt.Printf("Delete jobs by id successfully\n")
	return nil
}

func (m *JobInfoList)CateJobDisplay(version,project_name string) (joblist []JobList, err error) {

	sqlFmt := "SELECT\n    t1.job_name as job_name,\n    t1.status as status,\n       t1.pass_num as pass_num,\n       t1.fail_num as fail_num,\n       t1.exe_num as exe_num,\n       t1.debug_pending as debug_pending,\n      t1.tag as tag,\n  t1.source as source,\n       t1.comment as comment,\n       t1.owner as owner,\n       t1.log_url as log_url,\n       t1.finished_time as finished_time\nFROM\n    "+TNjobList()+" t1,\n    ( SELECT job_name, MAX( id ) maxId FROM "+TNjobList()+" GROUP BY job_name ) t2\nWHERE\n        t1.id = t2.maxId\n        and\n      t1.release_version='"+version+"'\n  AND t1.job_name IN ( SELECT job_name FROM "+TNjobInfoList()+" where project_name='"+project_name+"' )"
	sqlCount := "select count(*) cnt from "+TNjobInfoList()
	o := orm.NewOrm()
	var params []orm.Params

	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_,err = o.Raw(sqlFmt).QueryRows(&joblist)

	}
	fmt.Println(joblist)

	return
}