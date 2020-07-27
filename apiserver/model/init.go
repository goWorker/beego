package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self	*gorm.DB
	Docker 	*gorm.DB

}
var DB *Database

func (db *Database)Init(){
	DB = &Database{
		Self: GetSelfDB(),
		Docker: GetDockerDB(),
	}
}
func (db *Database) Close(){
	DB.Self.Close()
	DB.Docker.Close()
}
func openDB(username,password,addr,name string) *gorm.DB{
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
		)
	db,err := gorm.Open("mysql",config)
	if err != nil {
		log.Errorf(err,"Database connection failed. Database name: %s",name)
	}
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB){
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}

func InitSelfDB() *gorm.DB{
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB{
	return InitSelfDB()
}
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

