package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	// 添加注册的router
	beego.Router("/register", &controllers.RegisterController{})
	// 添加登录的router
	beego.Router("/login", &controllers.LoginController{})
	// 添加退出的router
	beego.Router("/exit", &controllers.ExitController{})
	// 添加文章路由
	beego.Router("/article/add", &controllers.AddArticleController{})
	// 显示文章内容页面
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	// 修改文章内容页面
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	// 删除文章页面
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	// 标签访问页面
	beego.Router("/tags", &controllers.TagsController{})
	// 访问相册操作
	beego.Router("/album", &controllers.AlbumController{})
	// 上传操作
	beego.Router("/upload", &controllers.UploadController{})
	// 关于我的页面
	beego.Router("/aboutme", &controllers.AboutMeController{})
}
