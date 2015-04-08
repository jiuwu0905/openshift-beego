package controllers

import (
	"beego/models"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {

	c.Ctx.WriteString(models.GetAll())
}
