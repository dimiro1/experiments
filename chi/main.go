package main

import (
	"net/http"

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

type Todo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

type TodoRenderer struct{ Todo }

func (t TodoRenderer) Render(w http.ResponseWriter, r *http.Request) error { return nil }

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
		todos := []Todo{
			{
				ID:    1,
				Title: "Example",
				Done:  false,
			},
		}
		w.Header().Set("Link", "<http://www.google.com>; rel=\"next\"")
		render.Respond(w, r, todos)
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo

		if err := render.Decode(r, &todo); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		render.Render(w, r, TodoRenderer{todo})
	})

	http.ListenAndServe(":7000", r)
}
