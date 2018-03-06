package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/dimiro1/experiments/custom-http-handler/handler"
	"github.com/dimiro1/experiments/custom-http-handler/imp"
	"github.com/dimiro1/experiments/custom-http-handler/render"
	"github.com/go-chi/chi"
)

type HelloWorldInput struct {
	Name string `schema:"name"`
}

func (h HelloWorldInput) IsValid() (bool, error) {
	if h.Name == "" {
		return false, errors.New("must cannot be blank")
	}

	return true, nil
}

type Index struct {
	render.Renderer

	// Caching?
	// Database?
	// Synchronizer?
	// Queue?
	// Logger?
	// Tracer?
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.Render(w, r, "Hello World")
}

type HelloWorld struct {
	handler.Default

	// Here you can add your own dependencies
	// Like database abstraction, external services abstractions
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input HelloWorldInput
	if err := h.Bind(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Render(w, r, "Bad Request")
		return
	}

	if _, err := h.Validate(input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Render(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Render(w, r, fmt.Sprintf("Hello %s", input.Name))
}

func main() {
	helloWorld := &HelloWorld{
		Default: handler.Default{
			Parameters: imp.Parameters{},
			Binder:     imp.ParametersBinder{},
			Renderer:   imp.Text{},
			Validator:  imp.Validator{},
		},
	}

	index := &Index{
		Renderer: imp.Text{},
	}

	r := chi.NewRouter()
	r.Get("/", index.ServeHTTP)
	r.Get("/hello/{name}", helloWorld.ServeHTTP)
	http.ListenAndServe(":8000", r)
}
