package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "iot-wrt.com"
	c.Data["Email"] = "147011391@qq.com"
	c.TplName = "index.html"
}
