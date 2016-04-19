package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Docker")
	})

	log.Println("Starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
