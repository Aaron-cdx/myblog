package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	Loginuser interface{}
}

// 这里负责session的处理，以及登录状态的设置
func (this *BaseController) Prepare(){
	loginuser := this.GetSession("loginuser")
	fmt.Println("login user ===>",loginuser)

	if loginuser != nil{
		this.IsLogin = true
		this.Loginuser = loginuser
	}else{
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}

