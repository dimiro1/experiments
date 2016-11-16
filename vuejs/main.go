// Complete example of usage of Vue.js with Go
// The idea of including vue templates in HTML came from Gitlab community project

package main

import "net/http"
import "html/template"

var templates = template.Must(
	template.New("").Delims("[[", "]]").ParseGlob("*.html"),
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", r.URL.Query())
	})

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":8080", nil)
}
