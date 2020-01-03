package controller

import (
	"fmt"
	"strconv"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/gin-gonic/gin"
)

// MemberController は Member の操作を行います
type MemberController struct {
}

// Index はメンバーの一覧を取得します
func (t *MemberController) Index(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")

	var Members []model.Member
	if err := db.Where("project_id = ?", pid).Find(&Members).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "member/index.html", gin.H{"Members": Members, "PID": pid})
}

// New はメンバーの新規作成画面に遷移します
func (m *MemberController) New(ctx *gin.Context) {
	ctx.HTML(200, "member/new.html", gin.H{"PID": ctx.Query("pid")})
}

// Create はメンバーの作成を行います
func (m *MemberController) Create(ctx *gin.Context) {
	db := db.GetDB()

	// TODO: バリデーション
	n := ctx.PostForm("name")
	e := ctx.PostForm("email")
	pid := ctx.PostForm("pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/member/index?pid="+pid)
	}

	member := model.Member{
		ProjectID: pidInt,
		Name:      n,
		Email:     e,
	}
	if err := db.Create(&member).Error; err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/member/new?pid="+pid)
	}

	ctx.Redirect(302, "/member/index?pid="+pid)
}
