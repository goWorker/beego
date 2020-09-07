package models


type JobInfoList struct {
	Id 				int 		`orm:"pk;auto;size(11)" json:"id"`
	ProjectName		string		`orm:"size(64)" json:"projectName"`
	ReleaseVersion 	string		`orm:"size(32)" json:"releaseVersion"`
	JobName			string		`orm:"size(64)" json:"jobName"`
}

func (m *JobInfoList)TableName() string {
	return TNjobInfoList()
}
