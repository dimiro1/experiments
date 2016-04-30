package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func() string {
		return "Hello World"
	})

	m.Get("/json", func(ctx *macaron.Context) {
		ctx.JSON(200, map[string]string{
			"hello": "world",
		})
	})

	m.Run()
}
