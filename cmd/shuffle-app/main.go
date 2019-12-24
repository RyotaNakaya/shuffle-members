package main

import (
	ctrl "github.com/RyotaNakaya/shuffle-members/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*.html")
	router.LoadHTMLGlob("web/template/project/*.html")

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{})
	// })

	p := router.Group("/project")
	{
		ctrl := ctrl.ProjectController{}
		p.GET("/index", ctrl.FecthAll)
		// p.GET("/:id", ctrl.Show)
		// p.POST("", ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	router.Run()
}
