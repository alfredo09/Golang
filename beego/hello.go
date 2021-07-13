package main

import (
	_ "beego_jwt_sample/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Register `mysql` driver
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// Register `default` database
	_ = orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/beego_jwt_test?charset=utf8")
	// Run migrations to create tables
	force := true // Drop old table and create new
	err := orm.RunSyncdb("default", force, beego.BConfig.RunMode == "dev")
	if err != nil {
		fmt.Printf("An Error Occurred: %v", err)
	}
	fmt.Printf("Conectado correctamente")
}

/* func main() {
	// Set configurations
	beego.BConfig.RunMode = "dev"

	beego.BConfig.AppName = "BeegoJWT"
	beego.BConfig.Listen.HTTPPort = 80

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.RouterCaseSensitive = false
	beego.BConfig.ServerName = "BEEGO_JWT"

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.Listen.HTTPPort = 8080

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.Listen.EnableAdmin = true
	}

	// Run application
	beego.Run()
} */
/* package main

import (
        "github.com/beego/beego/v2/server/web"
)

type MainController struct {
        web.Controller
}

func (this *MainController) Get() {
        this.Ctx.WriteString("hello world")   
}

func main() {
        web.Router("/", &MainController{})    
        web.Run()
} */