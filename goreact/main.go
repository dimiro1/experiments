package main

import (
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/olebedev/go-duktape"
)

func loadJSFile(ctx *duktape.Context, file string) error {
	ctx.EvalFile(file)
	result := ctx.GetString(-1)

	if result != "" {
		return errors.New(result)
	}

	ctx.Pop()

	return nil
}

func loadJSFiles(ctx *duktape.Context, files ...string) error {
	for _, file := range files {
		err := loadJSFile(ctx, file)

		if err != nil {
			return err
		}
	}

	return nil
}

func renderServer(ctx *duktape.Context, name string) (string, error) {
	ctx.GetGlobalString("renderServer")

	if ctx.IsUndefined(-1) {
		return "", errors.New("Could not find function 'renderServer'")
	}

	ctx.PushString(name)
	ctx.Call(1)
	result := ctx.GetString(-1)
	ctx.Pop()

	return result, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join("templates", "index.gohtml"))

	if err != nil {
		log.Fatal(err)
	}

	ctx := duktape.New()

	err = loadJSFiles(ctx,
		"static/duktape-polyfill.js",
		"static/react.js",
		"static/react-dom-server.js",
		"static/components.js",
		"static/server.js",
	)

	if err != nil {
		log.Fatal(err)
	}

	component, err := renderServer(ctx, "Claudemiro")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, component)
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
