package controller

import (
	"fmt"

	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ProjectController はプロジェクト操作します
type ProjectController struct {
}

// Show はプロジェクトを取得します
func (p *ProjectController) Show(ctx *gin.Context) {
	// DBから取ってくる
	err := errors.New("error")

	Project := model.Project{
		ID:          1,
		Name:        "プロジェクト",
		Description: "説明",
	}

	err = nil

	if err != nil {
		fmt.Println(err)
	} else {
		ctx.HTML(200, "show.html", Project)
	}
}

// FecthAll はプロジェクトの一覧を取得します
func (p *ProjectController) FecthAll(ctx *gin.Context) {
	// DBから取ってくる
	err := errors.New("error")

	Projects := []model.Project{
		model.Project{
			ID:          1,
			Name:        "プロジェクト",
			Description: "説明",
		},
		model.Project{
			ID:          2,
			Name:        "プロジェクト2",
			Description: "説明",
		},
	}

	err = nil

	if err != nil {
		fmt.Println(err)
	} else {
		ctx.HTML(200, "index.html", Projects)
	}
}

// Create はプロジェクトの作成を行います
func (p *ProjectController) Create() string {
	// value := c.QueryParam("value")
	return "Create"
}
