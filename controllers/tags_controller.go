/**
 * @Author: caoduanxi
 * @Date: 2021/3/17 14:46
 * @Motto: Keep thinking, keep coding!
 * 标签控制器
 */

package controllers

import "myblog/models"

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}
