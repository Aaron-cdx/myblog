/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 16:00
 * @Motto: Keep thinking, keep coding!
 * 上传图片模块控制器
 */

package controllers

import (
	"github.com/prometheus/common/log"
	"myblog/models"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		log.Error("find albums meet error...", err.Error())
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
