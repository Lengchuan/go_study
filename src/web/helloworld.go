package main

import "github.com/astaxie/beego"

func main() {

	beego.Router("/", &HomeController{})
	beego.Run()
}

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello world!!!")
}
