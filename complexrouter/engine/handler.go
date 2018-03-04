package engine

type Handler interface {
	Serve(ctx *Context) error
}

type HandlerFunc func(ctx *Context) error

func (h HandlerFunc) Serve(ctx *Context) error {
	return h(ctx)
}
