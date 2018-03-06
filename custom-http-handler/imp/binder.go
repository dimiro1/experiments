package imp

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type ParametersBinder struct{}

func (ParametersBinder) Bind(r *http.Request, dst interface{}) error {
	src := map[string][]string{}

	if rctx := chi.RouteContext(r.Context()); rctx != nil {
		for _, k := range rctx.URLParams.Keys {
			src[k] = append(src[k], rctx.URLParam(k))
		}
	}

	return decoder.Decode(dst, src)
}
