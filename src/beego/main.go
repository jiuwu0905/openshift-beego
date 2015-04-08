package main

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api", &controllers.ApiController{})
	beego.Run()
}
