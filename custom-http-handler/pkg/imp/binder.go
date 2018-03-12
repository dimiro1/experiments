package imp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type ParametersBinder struct{}

func (ParametersBinder) Bind(r *http.Request, dst interface{}) error {
	src := map[string][]string{}

	// Route parameters
	if routeContext := chi.RouteContext(r.Context()); routeContext != nil {
		for _, k := range routeContext.URLParams.Keys {
			src[k] = append(src[k], routeContext.URLParam(k))
		}
	}

	// Query
	for k, v := range r.URL.Query() {
		src[k] = append(src[k], v...)
	}

	return decoder.Decode(dst, src)
}

type JSONBinder struct{}

func (JSONBinder) Bind(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

type XMLBinder struct{}

func (XMLBinder) Bind(r *http.Request, dst interface{}) error {
	return xml.NewDecoder(r.Body).Decode(dst)
}

type ContentNegotiationBinder struct {
	JSONBinder       JSONBinder
	XMLBinder        XMLBinder
	ParametersBinder ParametersBinder
}

func (c ContentNegotiationBinder) Bind(r *http.Request, dst interface{}) error {
	if r.ContentLength == 0 {
		if r.Method == http.MethodGet || r.Method == http.MethodDelete {
			return c.ParametersBinder.Bind(r, dst)
		}
	}

	switch contentType(r) {
	case "xml":
		return c.XMLBinder.Bind(r, dst)
	case "json":
		fallthrough
	default:
		return c.JSONBinder.Bind(r, dst)
	}
}
