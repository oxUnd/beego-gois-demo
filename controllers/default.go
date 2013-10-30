package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xiangshouding/gofis"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println("Index")
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "demo/page/index.tpl"
}

func (this *MainController) Render() error {
	rb, err := this.RenderBytes()
	if err != nil {
		return err
	} else {
		this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")

		rb, _ = gofis.AfterProcess(rb)

		this.Ctx.Output.Body(rb)
	}
	return nil
}
