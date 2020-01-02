package controller

import (
	"fmt"
	"strconv"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/gin-gonic/gin"
)

// TagController は Tag の操作を行います
type TagController struct {
}

// Show はタグを取得します
func (t *TagController) Show(ctx *gin.Context) {
}

// Index はタグの一覧を取得します
func (t *TagController) Index(ctx *gin.Context) {
	db := db.GetDB()
	var Tags []model.Tag
	pid, err := strconv.Atoi(ctx.Param("pid"))
	if err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/tag/index/"+ctx.Param("pid"))
	}

	if err := db.Where("project_id = ?", pid).Find(&Tags).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "tag/index.html", gin.H{"Tags": Tags, "PID": pid})
}

// New はタグの新規作成画面に遷移します
func (t *TagController) New(ctx *gin.Context) {
	ctx.HTML(200, "tag/new.html", gin.H{"PID": ctx.Param("pid")})
}

// Create はタグの作成を行います
func (t *TagController) Create(ctx *gin.Context) {
	db := db.GetDB()

	// TODO: バリデーション
	n := ctx.PostForm("name")
	pid, err := strconv.Atoi(ctx.PostForm("pid"))
	if err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/tag/index/"+ctx.PostForm("pid"))
	}

	tag := model.Tag{
		Name:      n,
		ProjectID: pid,
	}
	if err := db.Create(&tag).Error; err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.HTML(500, "tag/new.html", gin.H{"Error": err})
	}

	fmt.Println("success")

	ctx.Redirect(302, "/tag/index/"+ctx.PostForm("pid"))
}

// Delete はタグの削除を行います
func (t *TagController) Delete(ctx *gin.Context) {
	// db := db.GetDB()

	// id, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	// TODO: エラーハンドリング
	// 	ctx.Redirect(302, "/tag/index")
	// }

	// var tag model.Tag
	// if err := db.Delete(&tag, id).Error; err != nil {
	// 	fmt.Println(err)
	// 	// TODO: エラーハンドリング
	// 	ctx.Redirect(302, "/tag/index")
	// }

	// ctx.Redirect(302, "/tag/index")
}
