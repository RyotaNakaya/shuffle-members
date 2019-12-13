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

// Fecth はプロジェクトを取得します
func (p *ProjectController) Fecth(ctx *gin.Context) {
	// DBから取ってくる
	// var s user.Service
	// p, err := s.GetAll()
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
		ctx.HTML(200, "index.html", Project)
	}
}

// FecthAll はプロジェクトの一覧を取得します
func (p *ProjectController) FecthAll() string {
	// DBから一覧を取ってくる
	return "FecthAll"
}

// Create はプロジェクトの作成を行います
func (p *ProjectController) Create() string {
	return "Create"
}
