package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"goji.io/pat"

	"goji.io"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", pat.Param(ctx, "name"))
}

func main() {
	router := goji.NewMux()
	router.HandleFunc(pat.Get("/"), index)
	router.HandleFuncC(pat.Get("/hello/:name"), hello)

	log.Println("Starting server on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
