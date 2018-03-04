package engine

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Get(route string, handler HandlerFunc) {
	s.router.GET(route, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		ctxParams := Params{}
		for _, p := range params {
			ctxParams = append(ctxParams, Param{Key: p.Key, Value: p.Value})
		}

		ctx := &Context{
			Params:   ctxParams,
			Request:  r,
			Response: w,
		}

		err := handler.Serve(ctx)

		if err != nil {
			// Have to deal with errors
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

func NewRouter() *Server {
	return &Server{
		router: httprouter.New(),
	}
}
