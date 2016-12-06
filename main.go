package main

import (
	"gorouter/logs"

	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		logs.Debug.Println("a log request")
		return c.String(http.StatusOK, "Hello logs")
	})
	logs.Warning.Println("logs start")
	e.Start(":1112")
}
