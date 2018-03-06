package imp

import (
	"net/http"

	"github.com/gin-gonic/gin/json"
)

type Text struct{}

func (Text) Render(w http.ResponseWriter, r *http.Request, i interface{}) error {
	w.Header().Set("content-type", "text/plain")
	_, err := w.Write([]byte(i.(string)))
	return err
}

type JSON struct{}

func (JSON) Render(w http.ResponseWriter, r *http.Request, i interface{}) error {
	js, err := json.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json")
	_, err = w.Write(js)
	return err
}
