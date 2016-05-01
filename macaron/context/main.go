package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

type context struct {
	*macaron.Context
	name string
}

type repository interface {
	All() []string
}

type dummy struct{}

func (d dummy) All() []string {
	return []string{"Go", "Macaron"}
}

func contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &context{
			Context: c,
			name:    "Claudemiro",
		}

		c.Map(ctx)
	}
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(contexter())
	m.MapTo(dummy{}, (*repository)(nil))

	m.Get("/", func(ctx *context) string {
		return ctx.name
	})

	m.Get("/repo", func(ctx *context, r repository) {
		ctx.JSON(http.StatusOK, r.All())
	})

	m.Run()
}
