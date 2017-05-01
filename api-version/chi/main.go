package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

type todo struct {
	ID    uint64 `xml:"id"    json:"id"`
	Title string `xml:"title" json:"title"`
	Done  bool   `xml:"done"  json:"done"`
}

func main() {
	router := chi.NewRouter()

	v1 := chi.NewRouter()
	v1.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		todo := todo{
			ID:    1,
			Title: "Test",
			Done:  false,
		}
		render.Respond(w, r, todo)
	})

	v2 := chi.NewRouter()
	v2.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		todo := todo{
			ID:    1,
			Title: "Test",
			Done:  true,
		}
		render.Respond(w, r, todo)
	})

	apis := map[string]*chi.Mux{
		"application/vnd.myapi.v1": v1,
		"application/vnd.myapi.v2": v2,
	}
	latestAPI := v2

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		version := r.Header.Get("Accept")

		api, ok := apis[version]
		if ok {
			api.ServeHTTP(w, r)
			return
		}

		latestAPI.ServeHTTP(w, r)
	})

	http.ListenAndServe(":9000", router)
}
