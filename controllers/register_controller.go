package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"myblog/models"
	"myblog/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

// 处理注册请求的业务逻辑
func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	// 获取判断是否已经存在
	id := models.QueryUserWithUsername(username)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		this.ServeJSON()
		return
	}

	// 获取存储的MD5密码值
	password = utils.MD5(password)
	user := models.User{Username: username, Password: password, Status: 0, CreateTime: time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户注册失败", "error": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "用户注册成功"}
	}
	this.ServeJSON()
}
