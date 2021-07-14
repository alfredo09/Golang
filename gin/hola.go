package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	
)

func getMysqlDB() (*sql.DB, error) {
	USER_NAME := "root"
	PASS_WORD := "root"
	HOST := "172.31.23.166"
	PORT := "3306"
	DATABASE := "mysql"
	CHARSET := "utf8"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET))
	if err != nil {
		return nil, err
	}
	return db, nil
}
// Inicializado enlace
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

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "holas soy io :v",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
} */