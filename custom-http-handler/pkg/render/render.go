package render

import "net/http"

type Renderer interface {
	Render(w http.ResponseWriter, r *http.Request, status int, data interface{}) error
}
