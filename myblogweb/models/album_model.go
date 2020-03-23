package models

import "myblogweb/utils"
type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}

func InsertAlbum(album Album) (int64, error) {
	return utils.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",
		album.Filepath, album.Filename, album.Status, album.Createtime)
}