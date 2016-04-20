package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		files, _ := filepath.Glob("*")
		json.NewEncoder(w).Encode(files)
	})

	log.Println("Starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
