package models

import (
	"bytes"
	"fmt"
	"html/template"
	"newblog/utils"
	"strconv"
	"strings"
)

type User struct {
	Id	int
	Username	string
	Password	string
	Status 		int
	Createtime	int64
}

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       [] TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

func InsertUser(user User)(int64, error){
	fmt.Println(user.Username, user.Password, user.Status, user.Createtime)
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}


//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	fmt.Println(row)
	id := 0
	row.Scan(&id)
	return id
}

func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username=%s", username)
	return QueryUserWightCon(sql)
}

func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

func  MakeHomeBlocks(articles []Article,isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		t,_ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer,homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->",htmlHome)
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string) []TagLink {
	var tagLink [] TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}