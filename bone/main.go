package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", bone.GetValue(r, "name"))
}

func main() {
	router := bone.New()

	router.GetFunc("/", index)
	router.GetFunc("/hello/:name", hello)

	log.Println("Starting server on port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
