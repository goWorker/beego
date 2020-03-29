package controllers

import (
	"fmt"
	"io"
	"newblog/models"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct{
	BaseController
}


func (this *UploadController)Post(){
	fmt.Println("File upload...")
	fileData, fileHeader,err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return
	}
	fmt.Println("name:",fileHeader.Filename,fileHeader.Size)
	fmt.Println(fileData)
	now := time.Now()

	fmt.Println("ext:",filepath.Ext(fileHeader.Filename))
	fileType := "other"

	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == "gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir,os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir,fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}
	_,err = io.Copy(desFile,fileData)
	if err != nil {
		this.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{0,filePathStr,fileName,0,timeStamp}
		models.InsertAlbum(album)
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	this.ServeJSON()
}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	this.ServeJSON()
}