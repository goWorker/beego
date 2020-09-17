package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lexkong/log"
	"time"
)

type JobAggregateList struct {
	Id 				int 		`orm:"pk;auto;size(11)" json:"id"`
	ProjectName    string		`orm:"size(100)" json:"project_name"`
	ReleaseVersion string		`orm:"size(10)" json:"release_version"`
	TotalJob  	int				`orm:"size(10)" json:"total_job"`
	Executed       int			`orm:"size(10)" json:"executed"`
	Pass           int			`orm:"size(10)" json:"pass"`
	Fail           int			`orm:"size(10)" json:"fail"`
	ExeRatio       float32		`orm:"size(10)" json:"exe_ratio"`
	PassRatio     float32		`orm:"size(10)" json:"pass_ratio"`
	FailRatio     float32		`orm:"size(10)" json:"fail_ratio"`
	SaveTime		time.Time	`orm:"default(auto_now_add);type(datetime)" json:"save_time"`
}
type SaveTimeStr struct {
	SaveTime		time.Time	`json:"save_time"`
}
type AggregateHistory struct{
	ProjectName    string		`json:"project_name"`
	ReleaseVersion string		`json:"release_version"`
	TotalJob  	int				`json:"total_job"`
	Executed       int			`json:"executed"`
	Pass           int			`json:"pass"`
	Fail           int			`json:"fail"`
	ExeRatio       float32		`json:"exe_ratio"`
	PassRatio     float32		`json:"pass_ratio"`
	FailRatio     float32		`json:"fail_ratio"`

}
type AggregateHistoryList struct {
	SaveTimeHis		time.Time	`json:"save_time"`
	AggregateHistoryLi 	[]AggregateHistory

}

func (m *JobAggregateList) TableName() string {
	return TNjobAggregateList()
}

func NewJobAggregateList() *JobAggregateList {
	return &JobAggregateList{}
}

func (m *JobAggregateList)SaveAggregateList(jobcatesummlist []JobCateSummList,version string) (jobaggregatelist []JobAggregateList){
	if len(jobcatesummlist) == 0 {
		return nil
	}else{
		var timeLayoutStr = "2006-01-02 15:04:05"
		t := time.Now().UTC()
		ts := t.Format(timeLayoutStr) //time convert tostring
		st, _ := time.Parse(timeLayoutStr, ts) //string convert to time
		fmt.Printf("The aggregate save time is: %v\n",st)
		var jobaggregatelist = make([]JobAggregateList,len(jobcatesummlist))
		for i,ele := range(jobcatesummlist){
			jobaggregatelist[i].ProjectName=ele.ProjectName
			jobaggregatelist[i].TotalJob = ele.TotalJob
			jobaggregatelist[i].Executed = ele.Executed
			jobaggregatelist[i].Pass = ele.Pass
			jobaggregatelist[i].Fail = ele.Fail
			jobaggregatelist[i].ExeRatio = ele.ExeRatio
			jobaggregatelist[i].PassRatio = ele.PassRatio
			jobaggregatelist[i].FailRatio = ele.FailRatio
			jobaggregatelist[i].ReleaseVersion = version
			jobaggregatelist[i].SaveTime = st
		}
		return jobaggregatelist
	}
}

func (m *JobAggregateList)CheckSaveTime(version string) (aggregatehistorylist []AggregateHistoryList, err error){
	var timeLayoutStr = "2006-01-02 15:04:05"
	var savetimelist = make([]SaveTimeStr,0)
	sqlfmt := "select distinct save_time from "+TNjobAggregateList()+" where release_version='"+version+"' order by save_time desc limit 100"
	o := orm.NewOrm()
	_,err = o.Raw(sqlfmt).QueryRows(&savetimelist)
	fmt.Printf("The save time list is:%v\n",savetimelist)
	if err != nil{
		return
	}
	if len(savetimelist) == 0 {
		return
	}else{
		var aggregatehistorylist = make([]AggregateHistoryList,len(savetimelist))

		for i,st := range(savetimelist){
			var aggregatehistory = make([]AggregateHistory,0)
			timeToStr:=st.SaveTime.Format(timeLayoutStr)
			fmt.Printf("The time string is: %s\n",timeToStr)
			sqlcheckgroup := "select project_name,release_version, total_job,executed,pass,fail,exe_ratio,pass_ratio,fail_ratio from jobaggregatelist where release_version='6.1.0' and save_time='"+timeToStr+"'"
			_,err = o.Raw(sqlcheckgroup).QueryRows(&aggregatehistory)
			if err != nil{
				log.Error("Encounter error: %v",err)
			}
			fmt.Printf("The aggregatehistory is: %v\n",aggregatehistory)
			aggregatehistorylist[i].SaveTimeHis = st.SaveTime
			aggregatehistorylist[i].AggregateHistoryLi = aggregatehistory
		}

		fmt.Printf("The aggregatehistorylist is: %v\n",aggregatehistorylist)
		return aggregatehistorylist,nil
	}

}