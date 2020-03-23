package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterModel(
		new(Category),
		new(Book),
		new(BookCategory),
		new(Member),
		new(Collection),
		new(Fans),
		)
}

func TNCategory()string{
	return "md_category"
}

func TNBook()string{
	return "md_books"
}

func TNBookCategory() string{
	return "md_book_category"
}

func TNMembers() string {
	return "md_members"
}

func TNCollection() string {
	return "md_star"
}

func TNFans() string {
	return "md_fans"
}

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
