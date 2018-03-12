package imp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Text struct{}

func (Text) Render(w http.ResponseWriter, r *http.Request, status int, i interface{}) error {
	w.Header().Set("Content-Type", "text/plain")
	var data []byte

	// Specific types
	switch i.(type) {
	case string:
		data = []byte(i.(string))
	case error:
		data = []byte(i.(error).Error())
	}

	// Stringer
	if s, ok := i.(fmt.Stringer); ok {
		data = []byte(s.String())
	}

	w.WriteHeader(status)
	_, err := w.Write(data)
	return err
}

type JSON struct{}

func (JSON) Render(w http.ResponseWriter, r *http.Request, status int, i interface{}) error {
	switch i.(type) {
	case error:
		i = struct {
			Message string `json:"message"`
		}{
			i.(error).Error(),
		}
	}

	js, err := json.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	return err
}

type XML struct{}

func (XML) Render(w http.ResponseWriter, r *http.Request, status int, i interface{}) error {
	switch i.(type) {
	case error:
		i = struct {
			XMLName xml.Name `xml:"error"`
			Message string   `xml:"message,attr"`
		}{
			Message: i.(error).Error(),
		}
	}

	x, err := xml.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(status)
	_, err = w.Write(x)
	return err
}

type ContentNegotiationRenderer struct {
	JSONRenderer JSON
	XMLRenderer  XML
	TextRenderer Text
}

func (c ContentNegotiationRenderer) Render(w http.ResponseWriter, r *http.Request, status int, i interface{}) error {
	switch contentType(r) {
	case "xml":
		return c.XMLRenderer.Render(w, r, status, i)
	case "json":
		return c.JSONRenderer.Render(w, r, status, i)
	default:
		return c.TextRenderer.Render(w, r, status, i)
	}
}
