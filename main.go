package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "myblog/routers"
	"myblog/utils"
)

///*type User struct {
//	Id   int
//	Name string `orm:"size(100)"`
//}
//
//func init() {
//	orm.RegisterDataBase("default", "mysql", "root:cao236476.@tcp(121.5.58.228:3306)/myblog?charset=utf8&loc=Local")
//
//	orm.RegisterModel(new(User))
//
//	orm.RunSyncdb("default", false, true)
//}*/
func init() {
	utils.InitMysql()
}

func main() {
	/**
	实现session启用可以通过配置文件设置
	sessionon = true
	或者在main方法中写入
	beego.BConfig.WebConfig.Session.SessionOn = true
	*/
	beego.Run()

}
