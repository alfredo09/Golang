import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "172.31.23.166"
	PORT      = "3306"
	DATABASE  = "mysql"
	CHARSET   = "utf8"
)

// Inicializado enlace
func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	// Abrir la conexión falló
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("La configuración de la fuente de datos es incorrecta:" + MysqlDbErr.Error())
	}
	// Número de conexión máxima
	MysqlDb.SetMaxOpenConns(100)
	// conexión inactiva
	MysqlDb.SetMaxIdleConns(20)
	// Ciclo de conexión máxima
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("Falló el enlace de la base de datos:" + MysqlDbErr.Error())
	}
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