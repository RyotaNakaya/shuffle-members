package controller

import (
	"fmt"

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

// FecthAll はプロジェクトの一覧を取得します
func (p *ProjectController) FecthAll(ctx *gin.Context) {
	db := db.GetDB()
	var Projects []model.Project

	if err := db.Find(&Projects).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(Projects)

	ctx.HTML(200, "index.html", Projects)
}

// Create はプロジェクトの作成を行います
func (p *ProjectController) Create(ctx *gin.Context) {
}
