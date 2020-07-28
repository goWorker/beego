package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(JobList),

	)
}

/*
* Table Names
* */

func TNcaseLog() string {
	return "caseLog"
}

func TNjob2case() string {
	return "job2case"
}
func TNjobList() string {
	return "joblist"
}

/*
* Tool Funcs
* */
//设置增减
//@param            table           需要处理的数据表
//@param            field           字段
//@param            condition       条件
//@param            incre           是否是增长值，true则增加，false则减少
//@param            step            增或减的步长
func IncOrDec(table string, field string, condition string, incre bool, step ...int) (err error) {
	mark := "-"
	if incre {
		mark = "+"
	}
	s := 1
	if len(step) > 0 {
		s = step[0]
	}
	sql := fmt.Sprintf("update %v set %v=%v%v%v where %v", table, field, field, mark, s, condition)
	_, err = orm.NewOrm().Raw(sql).Exec()
	return
}

