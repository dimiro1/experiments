package main

import "net/http"

func main() {
	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger.json")
	})
	http.Handle("/doc/", http.StripPrefix("/doc/", http.FileServer(http.Dir("swagger"))))
	http.ListenAndServe(":9000", nil)
}
