package engine

import (
	"net/http"

	"github.com/gin-gonic/gin/json"
)

type M map[string]interface{}

type Context struct {
	Params   Params
	Request  *http.Request
	Response http.ResponseWriter
}

func (ctx *Context) Text(status int, text string) error {
	ctx.Response.WriteHeader(status)
	ctx.Response.Header().Set("Content-Type", "text/plain")
	_, err := ctx.Response.Write([]byte(text))
	return err
}

func (ctx *Context) JSON(status int, data interface{}) error {
	ctx.Response.WriteHeader(status)
	ctx.Response.Header().Set("Content-Type", "application/json")

	marshaled, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = ctx.Response.Write(marshaled)
	return err
}
