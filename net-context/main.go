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

func main() {
	ctx := context.WithValue(context.Background(), "Name", "Application Name")

	http.Handle("/", appHandler{ctx: ctx, Handler: indexHandler})
	http.ListenAndServe(":8080", nil)
}

func indexHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := ctx.Value("Name").(string)

	fmt.Fprintf(w, "indexHandler: name is %s", name)
}
