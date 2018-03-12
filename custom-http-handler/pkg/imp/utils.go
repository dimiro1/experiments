package imp

import (
	"net/http"
	"strings"
)

func contentType(r *http.Request) string {
	switch {
	case strings.HasPrefix(r.Header.Get("Accept"), "text/xml"),
		strings.HasPrefix(r.Header.Get("Content-Type"), "text/xml"):
		return "xml"
	case strings.HasPrefix(r.Header.Get("Accept"),
		"application/json"), strings.HasPrefix(r.Header.Get("Content-Type"), "application/json"):
		return "json"
	}

	return "text"
}
