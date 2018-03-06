package params

import "net/http"

type Parameters interface {
	ByName(*http.Request, string) string
}
