package main

import (
	"errors"
	"fmt"
	"net/http"

	"time"

	"github.com/dimiro1/experiments/custom-http-handler/pkg/cache"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/handler"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/imp"
	"github.com/dimiro1/experiments/custom-http-handler/pkg/render"
	"github.com/go-chi/chi"
)

type HelloWorldInput struct {
	Name string `schema:"name"`
}

func (h HelloWorldInput) IsValid() (bool, error) {
	if h.Name == "" {
		return false, errors.New("cannot be blank")
	}

	return true, nil
}

type Index struct {
	render.Renderer
	cache.Cache

	// Database?
	// Synchronizer?
	// Queue?
	// Logger?
	// Tracer?
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := i.Cache.Set("greeting", "Hello World", 2*time.Second)
	if err != nil {
		i.Render(w, r, err)
		return
	}

	cachedString, found := i.Cache.Get("greeting")
	if !found {
		i.Render(w, r, err)
		return
	}

	i.Render(w, r, cachedString.(string))
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
		Cache:    imp.NewCache(),
	}

	r := chi.NewRouter()
	r.Get("/", index.ServeHTTP)
	r.Get("/hello/{name}", helloWorld.ServeHTTP)
	http.ListenAndServe(":8000", r)
}
