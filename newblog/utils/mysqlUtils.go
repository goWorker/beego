package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql(){
	fmt.Println("Init MySQL...")
	driverName := beego.AppConfig.String("driverName")
	orm.RegisterDriver(driverName,orm.DRMySQL)

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")

	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db, _=sql.Open(driverName,dbConn)
	CreateTableWithUser()
	CreateTableWithArticle()
	CreateTableWithAlbum()
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}

func CreateTableWithArticle(){
	sql:=`create table if not exists article(
        id int(4) primary key auto_increment not null,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        createtime int(10)
        );`
	ModifyDB(sql)
}

func CreateTableWithAlbum() {
	sql := `create table if not exists album(
        id int(4) primary key auto_increment not null,
        filepath varchar(255),
        filename varchar(64),
        status int(4),
        createtime int(10)
        );`
	ModifyDB(sql)
}