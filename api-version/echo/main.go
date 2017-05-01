package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	v1 := echo.New()
	v1.GET("/hello", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello From v1")
	})

	v2 := echo.New()
	v2.GET("/hello", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello From v2")
	})

	apis := map[string]*echo.Echo{
		"application/vnd.myapi.v1": v1,
		"application/vnd.myapi.v2": v2,
	}
	latestAPI := v2

	router.Any("/*", func(ctx echo.Context) error {
		req := ctx.Request()
		res := ctx.Response()
		version := req.Header.Get("Accept")

		api, ok := apis[version]
		if ok {
			api.ServeHTTP(res, req)
			return nil
		}

		latestAPI.ServeHTTP(res, req)
		return nil
	})

	router.Start(":9000")
}
