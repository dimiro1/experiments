package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type repository interface {
	All() []string
}

type dummy struct{}

func (d dummy) All() []string {
	return []string{"Go", "Echo"}
}

func repositoryMiddleware(r repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("repository", r)
			return next(ctx)
		}
	}
}

func main() {
	m := echo.New()
	m.Use(middleware.Logger())
	m.Use(middleware.Recover())
	m.Use(repositoryMiddleware(dummy{}))

	m.Get("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	m.Get("/json", func(ctx echo.Context) error {
		repo := ctx.Get("repository").(repository)

		return ctx.JSON(200, repo.All())
	})

	m.Run(standard.New(":9090"))
}
