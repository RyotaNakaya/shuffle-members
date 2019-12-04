package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.GET("/projects", func(ctx *gin.Context) {
		ctx.HTML(200, "/project/index.html", gin.H{"data": "todo"})
	})

	router.Run()
}
