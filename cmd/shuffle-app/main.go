package main

import (
	"github.com/RyotaNakaya/shuffle-members/db"
	ctrl "github.com/RyotaNakaya/shuffle-members/internal/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db.Init()
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("web/template/*.html")
	router.LoadHTMLGlob("web/template/project/*.html")

	p := router.Group("/project")
	{
		ctrl := ctrl.ProjectController{}
		p.GET("/index", ctrl.Index)
		p.GET("/new", ctrl.New)
		p.POST("/create", ctrl.Create)
		// p.GET("/:id", ctrl.Show)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	router.Run()
}
