package main

import (
	"net/http"

	"github.com/dimiro1/experiments/chi/entities"
	"github.com/dimiro1/experiments/chi/v1"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

func init() {
	render.Respond = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if _, ok := v.(error); ok {
			if _, ok := r.Context().Value(render.StatusCtxKey).(int); !ok {
				w.WriteHeader(500)
			}

			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		render.DefaultResponder(w, r, v)
	}
}

// ErrResponse our standard error response
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render sets http status code on response
func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest creates a new Invalid request error
func ErrInvalidRequest(err error) render.Renderer {
	return ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic")
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := []entities.Todo{
			{
				ID:    1,
				Title: "Example",
				Done:  false,
			},
		}

		render.Render(w, r, v1.NewTodosResponse(todos))
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		var todo entities.Todo

		if err := render.Decode(r, &todo); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		render.Render(w, r, v1.NewTodoResponse(todo))
	})

	http.ListenAndServe(":7000", r)
}
