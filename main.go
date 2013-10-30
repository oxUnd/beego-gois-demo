package main

import (
	"app/firstApp/controllers"
	"github.com/astaxie/beego"
	"github.com/xiangshouding/gofis"
	"path/filepath"
)

func main() {
	var settings = make(map[string]string)
	root, _ := filepath.Abs(beego.AppPath)
	settings["config_dir"] = root + "/" + beego.ViewsPath + "/config/"
	gofis.Register(settings)
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
