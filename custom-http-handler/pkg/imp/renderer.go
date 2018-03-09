package imp

import (
	"net/http"

	"github.com/gin-gonic/gin/json"
)

type Text struct{}

func (Text) Render(w http.ResponseWriter, r *http.Request, i interface{}) error {
	w.Header().Set("content-type", "text/plain")
	var data []byte

	switch i.(type) {
	case string:
		data = []byte(i.(string))
	case error:
		data = []byte(i.(error).Error())
	}

	_, err := w.Write(data)
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
