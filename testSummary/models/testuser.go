package models

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func PrintUsers() {
	db,err := sql.Open("mysql","root:wandl123@tcp(127.0.0.1:3306)/testsummary?charset=utf8")
	defer db.Close()
	if err != nil{
		return
	}
	stmt,err := db.Prepare("select * from testtable limit 10")
	defer stmt.Close()
	if err != nil {
		return
	}
	rows,err :=stmt.Query()

	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next(){
		var id int
		var name string
		rows.Scan(&id,&name)
		fmt.Println(id,name)
	}
}