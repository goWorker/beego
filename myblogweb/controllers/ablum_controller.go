package controllers

import (
	"log"
	"myblogweb/models"
)
type AlbumController struct{
	BaseController
}

func (this *AlbumController) Get()  {
	albums,err := models.FindAllAlbums()
	if err !=nil{
		log.Println(err)
	}
	this.Data["Album"] = albums
	this.TplName="album.html"
}