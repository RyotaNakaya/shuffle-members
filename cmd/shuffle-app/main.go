package main

import (
	ctrl "github.com/RyotaNakaya/shuffle-members/internal/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/template/*.html")
	router.LoadHTMLGlob("web/template/project/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	// router.GET("/projects", func(ctx *gin.Context) {
	// 	s := ctrl.Project{}
	// 	res := s.Create()
	// 	ctx.HTML(200, "/project/index.html", gin.H{"data": res})
	// })

	p := router.Group("/project")
	{
		ctrl := ctrl.ProjectController{}
		p.GET("", ctrl.Fecth)
		// p.GET("/:id", ctrl.Show)
		// p.POST("", ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	router.Run()
}
