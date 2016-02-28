package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/olebedev/go-duktape"
)

func readFile(file string) string {
	js, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(js)
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join("templates", "index.gohtml"))

	if err != nil {
		log.Fatal(err)
	}

	ctx := duktape.New()

	// Loading Javascript
	ctx.EvalString(readFile("static/duktape-polyfill.js"))
	ctx.EvalString(readFile("static/react.js"))
	ctx.EvalString(readFile("static/react-dom-server.js"))
	ctx.EvalString(readFile("static/components.js"))
	ctx.EvalString(readFile("static/server.js"))

	// Calling function renderServer
	ctx.GetGlobalString("renderServer")
	ctx.PushString("Claudemiro")
	ctx.Call(1)

	component := ctx.GetString(-1)

	t.Execute(w, component)
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
