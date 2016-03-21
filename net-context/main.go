package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

type appHandler struct {
	ctx     context.Context
	Handler func(context.Context, http.ResponseWriter, *http.Request)
}

// ServeHTTP could have error handling, authentication
func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(h.ctx, w, r)
}

func param(ctx context.Context, name string) string {
	return ctx.Value(name).(string)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Name", "Application Name")
	ctx = context.WithValue(ctx, "Version", "1.0.0")

	http.Handle("/", appHandler{ctx: ctx, Handler: indexHandler})
	http.ListenAndServe(":8080", nil)
}

func indexHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := param(ctx, "Name")
	version := param(ctx, "Version")

	fmt.Fprintf(w, "indexHandler: name is %s and version is %s", name, version)
}
