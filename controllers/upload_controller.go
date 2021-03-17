/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 16:01
 * @Motto: Keep thinking, keep coding!
 * 文件上传控制器
 */

package controllers

import (
	"fmt"
	"io"
	"myblog/models"
	"os"
	"path/filepath"
	"time"
)

type UploadController struct {
	BaseController
}

/**
文件上传控制器开发
*/
func (this *UploadController) Post() {
	fileData, fileHeader, err := this.GetFile("upload")
	if err != nil {
		//log.Println("upload file error...", err.Error())
		this.responseErr(err)
		return
	}

	// 获取当前时间
	now := time.Now()
	fileType := "other"
	// 返回文件的后缀
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	// 获取文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	// 构建当前文件夹
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		//log.Println("make dir error...", err.Error())
		this.responseErr(err)
		return
	}
	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	// 通过/分隔
	filePathStr := filepath.Join(fileDir, fileName)
	destFile, err := os.Create(filePathStr) // 构建当前的文件
	if err != nil {
		//log.Println("create file path file error...", err.Error())
		this.responseErr(err)
		return
	}
	// 开始对图片数据进行拷贝,实际就是写文件
	_, err = io.Copy(destFile, fileData)
	if err != nil {
		//log.Println("copy file data to dest file error...", err.Error())
		this.responseErr(err)
		return
	}
	if fileType == "img" {
		// 这里对文件的路径进行插入操作
		album := models.Album{
			Id:         0,
			Filepath:   filePathStr,
			Filename:   fileName,
			Status:     0,
			CreateTime: timeStamp,
		}
		models.InsertAlbum(album)
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "图片上传成功"}
	this.ServeJSON()
}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
	this.ServeJSON()
}
