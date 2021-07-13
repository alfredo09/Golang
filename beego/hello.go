package main
import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getMysqlDB() (*sql.DB, error) {
	connectString := "root:root@tcp(172.31.23.166:3306)/bdEcho"

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func main() {
	db, err := getMysqlDB()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Printf("Conectado correctamente")
} 
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