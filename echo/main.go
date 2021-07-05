package main

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
}