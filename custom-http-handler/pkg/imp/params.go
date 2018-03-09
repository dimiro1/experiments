package imp

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Parameters struct{}

func (p Parameters) ByName(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
