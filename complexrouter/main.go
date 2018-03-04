package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimiro1/experiments/complexrouter/engine"
)

func main() {
	r := engine.NewRouter()

	r.Get("/", func(ctx *engine.Context) error {
		return ctx.Text(http.StatusOK, "Hello World")
	})

	r.Get("/hello/:name", func(ctx *engine.Context) error {
		return ctx.JSON(http.StatusOK, engine.M{
			"greeting": fmt.Sprintf("Hello %s", ctx.Params.ByName("name")),
		})
	})

	log.Println("Starting on port 3000...")
	http.ListenAndServe(":3000", r)
}
