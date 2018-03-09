package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimiro1/experiments/complexrouter/engine"
)

type (
	Renderer interface {
		JSONRenderer
		TextRenderer
	}

	JSONRenderer interface {
		JSON(int, interface{}) error
	}

	TextRenderer interface {
		Text(int, string) error
	}

	Base struct {
		// Binder
		// Renderer
		// Params
		// Validator
		Renderer Renderer
	}
)

type Handler struct {
	Base

	// Logger
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Renderer.Text(http.StatusOK, "Hello World")
}

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
