package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func async(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, http.StatusAccepted)

		go func() {
			log.Print("Starting async...")
			h.ServeHTTP(httptest.NewRecorder(), r)
			log.Print("Finishing async.")
		}()
	}
}

func main() {
	http.HandleFunc("/async", async(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	}))

	http.ListenAndServe(":8080", nil)
}
