package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"officialsummary/utils"
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

	sqlFmt := "SELECT a.* FROM "+ TNjobList()+" a,( SELECT job_name, release_version, MAX( finished_time ) ftime FROM jobList GROUP BY job_name, release_version ) b WHERE a.job_name = b.job_name AND a.release_version = b.release_version AND a.finished_time = b.ftime AND a.release_version = '"+version+"'"
	sqlCount := "select count(*) cnt from "+TNjobList()
	fmt.Println(sqlFmt)
	//fmt.Println(sqlCount)
	o := orm.NewOrm()
	var params []orm.Params

	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		_, err = o.Raw(sqlFmt).QueryRows(&joblist)

	}
	fmt.Println(joblist)

	return
}

func (m *JobList)SearchJob(version,id string) (joblist []JobList, err error) {

	sqlFmt := "SELECT a.* FROM "+ TNjobList()+" a,( SELECT job_name, release_version, MAX( finished_time ) ftime FROM jobList GROUP BY job_name, release_version ) b WHERE a.job_name = b.job_name AND a.release_version = b.release_version AND a.finished_time = b.ftime AND a.release_version = '"+version+"' AND a.id = '"+id+"'"
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
func (m *JobList)SearchHistoryJobDetail(Id string)(joblist []JobList, err error){
	sqlFmt := "SELECT * FROM "+ TNjobList()+" where id =  '"+Id+"'"
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
func (m *JobList)SearchJobHistory(version,name string) (joblist []JobList, err error) {
	//select * from jobList where release_version = "6.1.0" and job_name = "610Felix_SRBWsizingAndContainerLSP";
	sqlFmt := "SELECT * FROM "+ TNjobList()+" where release_version =  '"+version+"'  and job_name = '"+name+"'"
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

func ModifyJob(joblist JobList) (int64, error) {

	return utils.ModifyDB("update jobList set status=?,pass_num=?,fail_num=?,exe_num=? debug_pending=? tag=? comment=? owner=? log_url=? where id=?",
		joblist.Status, joblist.PassNum,joblist.FailNum,joblist.ExeNum,joblist.DebugPending,joblist.Tag,joblist.Comment,joblist.Owner,joblist.LogUrl, joblist.Id)
}

func (joblist *JobList) InsertOrUpdate(cols ...string) (id int64, err error) {
	o := orm.NewOrm()
	id = int64(joblist.Id)

	//m.DocumentName = strings.TrimSpace(m.DocumentName)

	_, err = o.Update(joblist, cols...)
	return

}

func (m *JobList) SelectByDocId(id int) (job *JobList, err error) {
	if id <= 0 {
		return m, errors.New("Invalid parameter")
	}

	o := orm.NewOrm()
	err = o.QueryTable(m.TableName()).Filter("Id", id).One(m)
	if err == orm.ErrNoRows {
		return m, errors.New("数据不存在")
	}

	return m, nil
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
//func (joblist *JobList)DeleteJob(version,jobname string){
//	o := orm.NewOrm()
//	o.Begin()
//	_, err := o.Raw("delete from "+TNjobList() +" where release_version = "+ version +" and job_name = (select job_name from (select job_name from jobList where id="+id+") tt);").Exec()
//
//	if err != nil {
//		o.Rollback()
//	}else {
//		o.Commit()
//	}
//}

func (joblist *JobList)Delete(version,jobname string) error{
	o := orm.NewOrm()
	//modelStore := new(JobList)

	//if doc, err := m.SelectByDocId(docId); err == nil {
	//	o.Delete(doc)
	//	modelStore.Delete(docId)
	//}

	var jobs []*JobList

	_, err := o.QueryTable(joblist.TableName()).Filter("job_name",jobname).Filter("release_version",version ).All(&jobs)
	if err != nil {
		return err
	}

	for _, item := range jobs {
		jobId := item.Id
		o.QueryTable(joblist.TableName()).Filter("id", jobId).Delete()
		////删除document_store表对应的文档
		//modelStore.Delete(docId)
		//m.Delete(docId)
	}
	return nil
}

func (joblist *JobList)DeleteHisJob(id string) error{
	o := orm.NewOrm()
	//modelStore := new(JobList)

	//if doc, err := m.SelectByDocId(docId); err == nil {
	//	o.Delete(doc)
	//	modelStore.Delete(docId)
	//}

	var jobs []*JobList

	_, err := o.QueryTable(joblist.TableName()).Filter("id",id).All(&jobs)
	if err != nil {
		return err
	}

	for _, item := range jobs {
		jobId := item.Id
		o.QueryTable(joblist.TableName()).Filter("id", jobId).Delete()
		////删除document_store表对应的文档
		//modelStore.Delete(docId)
		//m.Delete(docId)
	}
	return nil
}