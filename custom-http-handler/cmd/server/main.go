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
	Age  int    `schema:"age"`
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
		i.Render(w, r, http.StatusInternalServerError, err)
		return
	}

	cachedString, found := i.Cache.Get("greeting")
	if !found {
		i.Render(w, r, http.StatusInternalServerError, err)
		return
	}

	i.Render(w, r, http.StatusOK, cachedString.(string))
}

type HelloWorld struct {
	handler.Default

	// Here you can add your own dependencies
	// Like database abstraction, external services abstractions
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input HelloWorldInput
	if err := h.Bind(r, &input); err != nil {
		h.Render(w, r, http.StatusBadRequest, err)
		return
	}

	if _, err := h.Validate(input); err != nil {
		h.Render(w, r, http.StatusBadRequest, err)
		return
	}

	h.Render(w, r, http.StatusOK, fmt.Sprintf("Hello %s, age %d", input.Name, input.Age))
}

func main() {
	// These initializations must be made by function constructors
	// This is just for demonstration purposes
	helloWorld := &HelloWorld{
		Default: handler.Default{
			Parameters: imp.Parameters{},
			Binder:     imp.ParametersBinder{},
			Renderer: imp.ContentNegotiationRenderer{
				JSONRenderer: imp.JSON{},
				XMLRenderer:  imp.XML{},
				TextRenderer: imp.Text{},
			},
			Validator: imp.Validator{},
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
