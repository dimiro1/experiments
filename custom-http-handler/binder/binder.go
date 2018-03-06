package binder

import "net/http"

type Binder interface{
	Bind(r *http.Request, dst interface{}) error
}
