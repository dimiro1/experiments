package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type item struct {
	Title string `json:"title"`
}

type store interface {
	All() []item
}

type staticStore struct{}

func (s staticStore) All() []item {
	return []item{
		{
			"Hello",
		},
		{
			"World",
		},
	}
}

type httpStore struct {
	URL string
}

func (s httpStore) All() []item {
	resp, err := http.Get(s.URL + "/all")
	if err != nil {
		return []item{}
	}
	defer resp.Body.Close()

	items := []item{}
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return []item{}
	}

	return items
}

func main() {
	staticStore := staticStore{}
	httpStore := httpStore{"http://localhost:9000"}
	tmp := template.Must(template.New("index").Parse(`
			<html>
				<body>
					<ul>
						{{range .}}<li>{{.Title}}</li>{{end}}
					</ul>
				</body>
			</html>
		`))

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(staticStore.All())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmp.Execute(w, httpStore.All())
	})

	http.ListenAndServe(":9000", nil)
}
