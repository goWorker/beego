package controllers

import (
	"fmt"
	"newblog/models"
)

type AlbumController struct{
	BaseController
}

func (this *AlbumController) Get(){

	albums,err := models.FindAllAlbums()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("album model: ",albums)
	this.Data["Album"] = albums
	this.TplName="album.html"
}