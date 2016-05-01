package main

import "github.com/gin-gonic/gin"

func main() {
	m := gin.Default()

	m.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	m.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"hello": "world",
		})
	})

	m.Run()
}
