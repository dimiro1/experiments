package render

import "net/http"

type Renderer interface {
	Render(http.ResponseWriter, *http.Request, interface{}) error
}
