package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallat/golang"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang")
	})
	e.GET("/caesar/:text", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"encypted": golang.Caesar(c.Param("text"), 4)})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
