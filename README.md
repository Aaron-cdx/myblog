# Go-100-05-Beego学习

## 1. 安装beego

```go
// 下载beego的安装包
go get -u github.com/beego/beego/v2@v2.0.0
// 可能会与遇到错误,如下图所示，然后开启set GO111MODULE=on即可，go env可以看环境变量配置，mac/Linux使用export GO111MODULE=on即可
set GO111MODULE=on
```

![安装beego错误](https://gitee.com/Caoduanxi/picture/raw/master/2021/3/3-15-15-41-39.jpg)

如果安装还是没有反应

```go
set GO111MODULE=on
set GOPROXY=https://goproxy.io

// 然后再执行,即可完成安装beego和bee
$ go get -u github.com/beego/beego/v2
$ go get -u github.com/beego/bee/v2
```

## 2. Beego简介

### 2.1 beego是什么

> Beego是一个使用Go语言开发的应用Web框架，框架开始于2012年，目的是为大家提供一个高效率的Web应用开发框架，该框架采用模块封装，使用简单，容易学习。对程序员来说，beego掌握起来非常简单，只需要关注业务逻辑实现即可，框架自动为项目需求提供不同的模块功能。

**特性**

* **简单化**：支持RESTful风格、MVC模型；可以使用bee工具类提高开发效率，比如监控代码修改进行热编译，自动化测试代码以及自动化打包部署等丰富的开发调试功能。
* **智能化**：beego框架封装了路由模块、支持智能路由、智能监控，并可以监控内存消耗，CPU使用以及goroutine的运行状况，方便开发者对线上应用进行监控分析。
* **模块化**：beego根据功能对代码进行了解耦封装，形成了Session、Cache、Log、配置解析、性能监控、上下文操作、ORM等独立的模块，方便开发者进行使用
* **高性能**：beego采用Go原生的http请求，goroutine的并发效率应付大流量的Web应用和API引用。

### 2.2 命令行工具Bee

**bee**

> bee是一个开发工具，协助Beego框架开发项目是进行创建项目、运行项目、热部署等相关的项目管理的工具，beego是源码负责开发、bee是工具负责构建和管理项目。

```go
USAGE
    bee command [arguments]

AVAILABLE COMMANDS

    version     Prints the current Bee version // 打印当前bee版本
    migrate     Runs database migrations	// 运行数据库的
    api         Creates a Beego API application // 构建一个beego的API应用
    bale        Transforms non-Go files to Go source files// 转义非go的文件到go的src中区
    fix         Fixes your application by making it compatible with newer versions of Beego
// 通过使得新版本的beego兼容来修复应用
    pro         Source code generator// 源代码生成器
    dev         Commands which used to help to develop beego and bee// 辅助开发beego和bee的
    dlv         Start a debugging session using Delve// 使用delve进行debbugging
    dockerize   Generates a Dockerfile for your Beego application // 为beego应用生成dockfile
    generate    Source code generator// 源代码生成器
    hprose      Creates an RPC application based on Hprose and Beego frameworks
    new         Creates a Beego application// 创建beego应用
    pack        Compresses a Beego application into a single file // 压缩beego项目文件
    rs          Run customized scripts// 运行自定义脚本
    run         Run the application by starting a local development server
// 通过启动本地开发服务器运行应用
    server      serving static content over HTTP on port// 通过HTTP在端口上提供静态内容
    update      Update Bee// 更新bee
```

```go
// 创建一个beego项目
bee new FirstBeego
// 运行beego项目
bee run
```

![项目启动的页面](https://gitee.com/Caoduanxi/picture/raw/master/2021/3/3-15-17-36-50.jpg)

## 3. Beego启动流程分析

### 3.1 程序入口

```go
import (
	_ "FirstBeego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

// -------------------routers-------------------
import (
	"FirstBeego/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {// 会先执行init()函数
    beego.Router("/", &controllers.MainController{})
}

// -------------------MainController-------------------
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
```

**Go语言执行顺序**

![Go语言代码执行顺序](https://gitee.com/Caoduanxi/picture/raw/master/2021/3/3-15-17-46-24.jpg)

**Beego的beego.Run()逻辑**

> 执行完init()方法之后，程序继续向下执行，到main函数，此时在main函数中执行`beego.Run()`，主要做了以下几件事：
>
> * 解析配置文件，即app.conf文件，获取其中的端口、应用名称等信息
> * 检查是否开启session，如果开启了session，会初始化一个session对象
> * 是否编译模板，beego框架会在项目启动的时候根据配置把views目录下的所有模板进行预编译，然后存放在map中，这样可以有效的提高模板运行的效率，不需要进行多次编译
> * 监听服务端口，根据app.conf文件配置端口，启动监听

## 4. Beego组织架构

**项目配置：conf**

**控制器：controllers**

> 该目录是存放控制器文件的目录，所谓控制器就是控制应用调用哪些业务逻辑，由controllers处理完HTTP请求以后，并负责返回给前端调用者。

**数据层：models**

> models层可以解释为实体层或者数据层，在models层中实现用户和业务数据的处理，主要和数据库表相关的一些操作会放在这个目录中实现，然后将执行后的结果数据返回给controller层。增删改查的操作都是在models中实现。

**路由层：routers**

> 路由层，即分发，对进来的后天的请求进行分发操作，当浏览器进行一个http请求达到后台的web项目的时候，必须要让程序能够根据浏览器的请求url进行不同的业务处理，从接受前端请求到判断执行具体的业务逻辑的过程的工作，就让routers来实现。

**静态资源目录：static**

> 在static目录下，存放的是web项目的静态资源文件，主要有css、img、js、html这几类文件。html中会存放应用的静态页面文件。

**视图模板：views**

> views中存放的就是应用存放html模板页面的目录，所谓模板，就是页面框架和布局是已经用html写好了的，只需要在进行访问和展示的时候，将获取到的数据动态填充到页面中，能够提高渲染效率。因此，模板文件是非常常见的一种方式。

整个项目的架构就是MVC的运行模式。

## 5. beego框架路由设置

在beego框架中，支持四种路由设置，分别是：**基础路由**、**固定路由**、**正则路由**和**自动路由**

**基础路由**

直接给过`beego.Get()`、`beego.Post()`、`beego.Put()`，`beego.Delete()`等方法进行路由的映射，。

```go
beego.Get("",func) // 表示Get
beego.Post("",func) // 表示Post
```

**固定路由**

```go
beego.Router("/",controller)
```

> Get请求就会对应到Get方法，Post对应到post方法，Delete对应到Delete方法，Header方法对应到Header方法。

**正则路由**

> 正则路由是指可以在进行固定路由的基础上，支持匹配一定格式的正则表达式，比如`:id`、`:username`自定义正则，file的路径和后缀切换以及全匹配等操作。

**自定义路由**

> 在开发的时候用固定匹配想要直接执行对应的逻辑控制方法，因此beego提供了可以自定义的自定义路由配置。

```go
beego.Router("/",&IndexController{},"")

// Router adds a patterned controller handler to BeeApp.
// it's an alias method of HttpServer.Router.
// usage:
//  simple router
//  beego.Router("/admin", &admin.UserController{})
//  beego.Router("/admin/index", &admin.ArticleController{})
//
//  regex router
//
//  beego.Router("/api/:id([0-9]+)", &controllers.RController{})
//
//  custom rules
//  beego.Router("/api/list",&RestController{},"*:ListFood")
//  beego.Router("/api/create",&RestController{},"post:CreateFood")
//  beego.Router("/api/update",&RestController{},"put:UpdateFood")
//  beego.Router("/api/delete",&RestController{},"delete:DeleteFood")
```

## 6. 静态文件的设置

在go的web项目中，一些静态资源文件，如果用户要访问静态资源文件，则我们也是能够访问到的，这需要我们的项目中进行静态资源设置。

```go
beego.SetStaticPath("/down1","download1")
```

这里的download目录是指的非go web项目的static目录下目录，而是开发者重新新建的另外的目录。

==2021-03-15==

## 7. Beego博客项目

beego的orm是可以自动创建表的，与python的django框架有的一拼。

在Go中Object类型的数据使用`interface{}`空的接口类型来代替。

**如果有js文件失效，注意清除缓存之后再来玩，否则添加的js不会生效。**

```go
// 首页显示内容,f
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	// for index, value := range objects{} 实现遍历
	for _, art := range articles {
		// 转换为模板所需要的数据
		homePageParam := HomeBlockParam{}
		homePageParam.Id = art.Id
		homePageParam.Title = art.Title
		homePageParam.Tags = createTagsLinks(art.Tags)
		homePageParam.Short = art.Short
		homePageParam.Content = art.Content
		homePageParam.Author = art.Author
		homePageParam.CreateTime = utils.SwitchTimeStampToData(art.CreateTime)
		homePageParam.Link = "/article/" + strconv.Itoa(art.Id)
		homePageParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homePageParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homePageParam.IsLogin = isLogin

		// 处理变量，利用ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, homePageParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome ===>", htmlHome)
	return template.HTML(htmlHome)
}
// 这里可以实现html模板的渲染和追加 最后以html代码的形式插入到具体的前端html展示页面
```

博客项目大概做了三天吧。就搞完了。基本的代码都是MVC结构，跟Java比较像，不过对HTML的支持，感觉beego做的更好一些。让人使用起来就很舒服的感觉。其他的就下面总结一下吧：

beego的项目目录结构如下：

![beego项目目录结构](https://gitee.com/Caoduanxi/picture/raw/master/2021/3/3-17-17-10-6.jpg)

负责和数据库交互的是model，model主要存放实体类和承接具体的数据请求等相关的方法操作，提供数据给controller层。

![beego整体结构](https://gitee.com/Caoduanxi/picture/raw/master/2021/3/3-17-17-11-6.jpg)

**路由的话主要有四种：**

* 默认路由：beego自带模块Post、Put、Delete、Head、Get等网络请求类型的对应方法
* 自动路由：自动实现映射到Post、Put、Delete、Get等
* 正则表达式路由：`"/article/:id"`接收参数的时候需要`idStr := this.Ctx.Input.Param(":id")`

* 自定义路由：在博客开发中基本就是自定义路由了`/article/add`

**Session的处理：**

* 配置文件中配置session相关的配置
* 代码中通过SessionConfig进行参数配置

**操作session**

* SetSession：设置session值
* GetSession：获取session值
* DelSession：删除session值

**View视图模板：**

* `controller.TplName`指定渲染当前页面的模板文件全称
* 模板文件中通过`{{.param}}`实现变量数据的获取操作
* `controller.Data["param"]=xxx`实现对页面的需要使用的变量进行赋值操作

## 参考

**博客项目的具体文档**
具体参考：[Go语言web框架Beego从入门到超神_千锋](https://www.bilibili.com/video/BV1cg4y1B7Ev)
> 有一说一，讲的不咋地，主要是看文档后面做的，他也主要看文档讲的
文档地址:
1. [项目搭建、登录注册、Session功能开发](https://github.com/Aaron-cdx/myblog/blob/master/day38_%E9%A1%B9%E7%9B%AE%E6%90%AD%E5%BB%BA%E3%80%81%E7%99%BB%E5%BD%95%E6%B3%A8%E5%86%8C%E5%92%8CSession%E5%8A%9F%E8%83%BD%E5%BC%80%E5%8F%91.md)
2. [写文章、项目首页、查看文章详情](https://github.com/Aaron-cdx/myblog/blob/master/day39_%E5%86%99%E6%96%87%E7%AB%A0%E3%80%81%E9%A1%B9%E7%9B%AE%E9%A6%96%E9%A1%B5%E5%92%8C%E6%9F%A5%E7%9C%8B%E6%96%87%E7%AB%A0%E8%AF%A6%E6%83%85%E5%8A%9F%E8%83%BD%E5%BC%80%E5%8F%91.md)
3. [修改文章、删除文章、标签](https://github.com/Aaron-cdx/myblog/blob/master/day40_%E4%BF%AE%E6%94%B9%E6%96%87%E7%AB%A0%E3%80%81%E5%88%A0%E9%99%A4%E6%96%87%E7%AB%A0%E5%92%8C%E6%96%87%E7%AB%A0%E6%A0%87%E7%AD%BE%E5%8A%9F%E8%83%BD%E5%BC%80%E5%8F%91.md)
4. [首页功能扩展、图片上传、关于、总结](https://github.com/Aaron-cdx/myblog/blob/master/day41_%E9%A6%96%E9%A1%B5%E5%8A%9F%E8%83%BD%E6%89%A9%E5%B1%95%E3%80%81%E5%9B%BE%E7%89%87%E4%B8%8A%E4%BC%A0%E5%92%8C%E5%85%B3%E4%BA%8E%E5%8A%9F%E8%83%BD%E5%BC%80%E5%8F%91.md)