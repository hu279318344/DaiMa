package main

import (
	"vscode/beeblog/controllers"
	"vscode/beeblog/models"
	_ "vscode/beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func databas() {
	//开启ORM调试模式
	orm.Debug = true
	//自助建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	//databas()
	//注册beego路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	//启动beego
	beego.Run()
}
