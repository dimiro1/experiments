package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type todo struct {
	ID    uint64 `xml:"id"    json:"id"`
	Title string `xml:"title" json:"title"`
	Done  bool   `xml:"done"  json:"done"`
}

func render(w http.ResponseWriter, r *http.Request, i interface{}) (err error) {
	var (
		accept = r.Header.Get("Accept")
		b      []byte
	)

	switch accept {
	case "application/xml":
		b, err = xml.Marshal(i)
	case "application/json":
		fallthrough
	default:
		b, err = json.Marshal(i)
	}

	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		todo := todo{
			ID:    1,
			Title: "Example",
			Done:  false,
		}

		if err := render(w, r, todo); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":9000", nil)
}
