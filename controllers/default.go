package controllers

import (
	"github.com/astaxie/beego"
	"jjj/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type LoginController struct{
	beego.Controller
}
func (this *LoginController) Get() {
	name := "走起"
	id,err := models.AddUser(name)
	if err != nil {
		this.Data["id"]= id
		this.TplName="index.tpl"
	}
} 