package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	m := echo.New()
	m.Use(middleware.Logger())
	m.Use(middleware.Recover())

	m.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	m.GET("/json", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{
			"hello": "world",
		})
	})

	m.Run(standard.New(":9090"))
}
