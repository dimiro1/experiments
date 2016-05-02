package main

import "github.com/gin-gonic/gin"

type repository interface {
	All() []string
}

type dummy struct{}

func (d dummy) All() []string {
	return []string{"Go", "Gin"}
}

func repositoryMiddleware(r repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("repository", r)
		ctx.Next()
	}
}

func main() {
	m := gin.Default()
	m.Use(repositoryMiddleware(dummy{}))

	m.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	m.GET("/json", func(ctx *gin.Context) {
		repo := ctx.MustGet("repository").(repository)

		ctx.JSON(200, repo.All())
	})

	m.Run()
}
