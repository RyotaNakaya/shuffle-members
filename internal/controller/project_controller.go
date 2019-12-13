package controller

import (
	"encoding/json"
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

	project := model.Project{}

	resJSON, err := json.Marshal(project)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		ctx.JSON(200, resJSON)
	}
}

// FecthAll はプロジェクトのの一覧を取得します
func (p *ProjectController) FecthAll() string {
	// DBから一覧を取ってくる
	return "FecthAll"
}

// Create はプロジェクトの作成を行います
func (p *ProjectController) Create() string {
	return "Create"
}
