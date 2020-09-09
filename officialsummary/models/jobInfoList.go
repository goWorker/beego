package models

import (
	"fmt"
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
func (m *JobInfoList)TableName() string {
	return TNjobInfoList()
}
func NewJobInfoList() *JobInfoList {
	return &JobInfoList{}
}
func (m *JobInfoList)CateForVersion(version string) (jobcatesummlist []JobCateSummList, err error) {

	sqlFmt := "SELECT t6.project_name as project_name,\n        t6.totalJob as total_job,\n        ifnull(t5.totalExe, 0) as executed,\n        ifnull(t5.totalPass, 0) as pass,\n        ifnull(t5.totalFail, 0) as fail,\n        ifnull(t5.totalExe, 0)/t6.totalJob as exe_ratio,\n        ifnull(t5.totalPass, 0)/t6.totalJob as pass_ratio,\n        ifnull(t5.totalFail, 0)/t6.totalJob as fail_ratio\n FROM (SELECT project_name,sum(exe_num) AS totalJob FROM jobinfolist a,\n                                                         (SELECT t1.job_name,t1.exe_num\n                                                          FROM jobList t1,\n                                                               (SELECT job_name, max(id) AS maxId FROM jobList WHERE release_version = '"+version+"' GROUP BY job_name) t2\n                                                          WHERE t1.id = t2.maxId) b\n       WHERE a.job_name = b.job_name GROUP BY project_name) t6\n          LEFT JOIN\n      (SELECT t4.project_name,sum(t3.exe_num) AS totalExe, sum(t3.pass_num) AS totalPass, sum(t3.fail_num) AS totalFail\n       FROM jobinfolist t4,\n            (SELECT t1.job_name,t1.exe_num,t1.pass_num,t1.fail_num FROM\n                                                                       jobList t1,\n                                                                       (SELECT job_name, max(id) AS maxId FROM jobList WHERE release_version = '"+version+"' AND DATE_FORMAT(finished_time, '%Y-%m-%d') >= DATE_FORMAT(DATE_SUB(NOW(), INTERVAL 1 DAY), '%Y-%m-%d') GROUP BY job_name) t2\n             WHERE t1.id = t2.maxId) t3\n       WHERE t4.job_name = t3.job_name  AND t4.release_version = '"+version+"' GROUP BY t4.project_name ) t5\n      ON t6.project_name = t5.project_name;"
	sqlCount := "select count(*) cnt from "+TNjobInfoList()
	o := orm.NewOrm()
	var params []orm.Params

	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_,err = o.Raw(sqlFmt).QueryRows(&jobcatesummlist)

	}
	fmt.Println(jobcatesummlist)

	return
}

func (m *JobInfoList) Insert() (err error) {
	if _, err = orm.NewOrm().Insert(m); err != nil {
		return
	}
	return err
}