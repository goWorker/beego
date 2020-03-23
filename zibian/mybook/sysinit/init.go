package sysinit

import (
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	dbinit()
	sysinit()
}