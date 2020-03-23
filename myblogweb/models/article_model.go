package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"myblogweb/utils"
	"strconv"
)

type Article struct {
	Id 		int
	Title 	string
	Tags 	string
	Short 	string
	Content 	string
	Author 	string
	Createtime 	int64
}



func AddArticle(article Article) (int64,error){
	i, err := insertArticle(article)
	return i, err
}

func insertArticle(article Article)(int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}


func FindArticleWithPage(page int) ([]Article,error) {
	num, _:= beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page,num)
}


func QueryArticleWithPage(page,num int) ([]Article,error) {
	sql := fmt.Sprintf("limit %d, %d",page*num,num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string)([]Article,error){
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList,nil
}

var artcileRowsNum = 0

func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}



//----------查询文章-------------

func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

func UpdateArticle(article Article) (int64,error) {
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

func DeleteArticle(artID int)(int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithArtId(artID int) (int64, error){
	return utils.ModifyDB("delete from article where id=?", artID)
}

func QueryArticleWithParam(param string) []string{
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article",param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next(){
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList,arg)
	}
	return paramList
}

func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticleWithCon(sql)
}

func FindAllAlbums() ([]Album, error) {
	rows, err := utils.QueryDB("select id,filepath,filename,status,createtime from album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createtime)
		album := Album{id, filepath, filename, status, createtime}
		albums = append(albums, album)
	}
	return albums, nil
}

