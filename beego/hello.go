package main

import "github.com/beego/beego/v2/server/web"

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
}
func main() {
	web.Run()
}
