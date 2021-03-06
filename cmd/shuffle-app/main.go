package main

import (
	"github.com/RyotaNakaya/shuffle-members/db"
	ctrl "github.com/RyotaNakaya/shuffle-members/internal/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// DB セットアップ
	db.Init()
	defer db.Close()

	router := gin.Default()
	// asset の読み込み
	router.Static("/public", "./public")
	// template の読み込み
	router.LoadHTMLGlob("web/template/**/*.html")

	// ルーティング
	router = setRouting(router)

	router.Run()
}

func setRouting(r *gin.Engine) *gin.Engine {
	c := ctrl.ProjectController{}
	r.GET("/", c.Index)

	p := r.Group("/project")
	{
		ctrl := ctrl.ProjectController{}
		p.GET("/index", ctrl.Index)
		p.GET("/new", ctrl.New)
		p.POST("/create", ctrl.Create)
		p.GET("/delete/:id", ctrl.Delete)
		p.GET("/show/:id", ctrl.Show)
	}

	t := r.Group("/tag")
	{
		ctrl := ctrl.TagController{}
		t.GET("/index", ctrl.Index)
		t.GET("/new", ctrl.New)
		t.POST("/create", ctrl.Create)
		t.GET("/delete/:id", ctrl.Delete)
		t.GET("/edit/:id", ctrl.Edit)
		t.POST("/update/:id", ctrl.Update)
	}

	m := r.Group("/member")
	{
		ctrl := ctrl.MemberController{}
		m.GET("/index", ctrl.Index)
		m.GET("/new", ctrl.New)
		m.POST("/create", ctrl.Create)
		m.GET("/delete/:id", ctrl.Delete)
		m.GET("/edit/:id", ctrl.Edit)
		m.POST("/update/:id", ctrl.Update)
	}

	s := r.Group("/shuffle")
	{
		ctrl := ctrl.ShuffleController{}
		s.GET("/index", ctrl.Index)
		s.POST("/shuffle", ctrl.Shuffle)
	}

	lm := r.Group("/log_management")
	{
		ctrl := ctrl.LogManagementController{}
		lm.GET("/index", ctrl.Index)
		lm.GET("/delete/:id", ctrl.Delete)
		lm.GET("/edit/:id", ctrl.Edit)
		lm.POST("/update/:id", ctrl.Update)
	}

	return r
}
