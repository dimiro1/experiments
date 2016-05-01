package main

import "gopkg.in/macaron.v1"

type context struct {
	*macaron.Context
	name string
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
	m.Use(macaron.Logger())
	m.Use(contexter())

	m.Get("/", func(ctx *context) string {
		return ctx.name
	})

	m.Run()
}
