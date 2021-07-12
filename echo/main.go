package main

import (
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "54.200.47.227"
	pass := "root"
	host := "tcp(54.200.47.227:3306)"
	nombreBaseDeDatos := "bdEcho"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := obtenerBaseDeDatos()
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
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola Mundo, estoy corriendo como Servidor!")
	})
	e.GET("/hola", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola Mundo, estoy corriendo como una API!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}  */