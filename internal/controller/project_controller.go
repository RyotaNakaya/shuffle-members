package controller

import (
	"fmt"
	"strconv"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/gin-gonic/gin"
)

// ProjectController は Project の操作を行います
type ProjectController struct {
}

// Show はプロジェクトを取得します
func (p *ProjectController) Show(ctx *gin.Context) {
}

// Index はプロジェクトの一覧を取得します
func (p *ProjectController) Index(ctx *gin.Context) {
	db := db.GetDB()
	var Projects []model.Project

	if err := db.Find(&Projects).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "index.html", Projects)
}

// New はプロジェクトの新規作成画面に遷移します
func (p *ProjectController) New(ctx *gin.Context) {
	ctx.HTML(200, "new.html", "")
}

// Create はプロジェクトの作成を行います
func (p *ProjectController) Create(ctx *gin.Context) {
	db := db.GetDB()

	n := ctx.PostForm("name")
	d := ctx.PostForm("description")
	// TODO: バリデーション

	prj := model.Project{
		Name:        n,
		Description: d,
	}
	if err := db.Create(&prj).Error; err != nil {
		fmt.Println(err)
		// TODO: エラーハンドリング
		ctx.HTML(500, "new.html", gin.H{"Error": err})
	}

	ctx.Redirect(302, "/project/index")
}

// Delete はプロジェクトの削除を行います
func (p *ProjectController) Delete(ctx *gin.Context) {
	db := db.GetDB()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		// TODO: エラーハンドリング
		ctx.Redirect(302, "/project/index")
	}

	var prj model.Project
	if err := db.Delete(&prj, id).Error; err != nil {
		fmt.Println(err)
		// TODO: エラーハンドリング
		ctx.Redirect(302, "/project/index")
	}

	ctx.Redirect(302, "/project/index")
}
