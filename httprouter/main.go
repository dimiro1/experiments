package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Index")
}

func hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", params.ByName("name"))
}

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/hello/:name", hello)

	log.Println("Starting server on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
