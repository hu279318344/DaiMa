package controllers

import (
	"vscode/beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	option := this.Input().Get("option")
	switch option {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "edit":
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

}