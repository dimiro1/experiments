//go:generate resources -declare -output=public_resources.go -var=Assets -trim=public/ -tag=embed public/**/*
package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(Assets))
	http.ListenAndServe(":9090", nil)
}
