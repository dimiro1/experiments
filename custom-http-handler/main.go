package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/dimiro1/experiments/custom-http-handler/handler"
	"github.com/dimiro1/experiments/custom-http-handler/imp"
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

type HelloWorld struct {
	handler.Default

	// Here you can add your own dependencies
	// Like database abstraction, external services abstractions
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var input HelloWorldInput
	if err := h.Binder.Bind(r, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Renderer.Render(w, r, "Bad Request")
		return
	}

	if _, err := h.Validator.Validate(input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Renderer.Render(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Renderer.Render(w, r, fmt.Sprintf("Hello %s", input.Name))
}

func main() {
	helloWorld := &HelloWorld{
		Default: handler.Default{
			Params:    imp.Parameters{},
			Binder:    imp.ParametersBinder{},
			Renderer:  imp.Text{},
			Validator: imp.Validator{},
		},
	}

	r := chi.NewRouter()
	r.Get("/{name}", helloWorld.ServeHTTP)
	http.ListenAndServe(":8000", r)
}
