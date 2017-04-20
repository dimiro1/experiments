package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type todo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	e.GET("/panic", func(ctx echo.Context) error {
		panic("Panic")
	})

	e.GET("/error", func(ctx echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError)
	})

	e.GET("/genericError", func(ctx echo.Context) error {
		return errors.New("Some error")
	})

	e.GET("/unauthorized", func(ctx echo.Context) error {
		return echo.ErrUnauthorized
	})

	e.GET("/todos", func(ctx echo.Context) error {
		todos := []todo{
			{
				ID:    1,
				Title: "Example",
				Done:  false,
			},
		}
		ctx.Response().Header().Set("Link", "<http://www.google.com>; rel=\"next\"")
		return ctx.JSON(http.StatusOK, todos)
	})

	e.GET("/map", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"Hello": "World",
			"One":   1,
		})
	})

	e.GET("/hello/:name", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, ctx.Param("name"))
	})

	e.Start(":9000")
}
