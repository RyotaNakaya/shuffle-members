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
	router.LoadHTMLGlob("web/template/**/*.html")

	c := ctrl.ProjectController{}
	router.GET("/", c.Index)

	p := router.Group("/project")
	{
		ctrl := ctrl.ProjectController{}
		p.GET("/index", ctrl.Index)
		p.GET("/new", ctrl.New)
		p.POST("/create", ctrl.Create)
		p.GET("/delete/:id", ctrl.Delete)
		p.GET("/show/:id", ctrl.Show)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	t := router.Group("/tag")
	{
		ctrl := ctrl.TagController{}
		t.GET("/index/:pid", ctrl.Index)
		t.GET("/new/:pid", ctrl.New)
		t.POST("/create", ctrl.Create)
		t.GET("/delete/:id", ctrl.Delete)
	}

	router.Run()
}
