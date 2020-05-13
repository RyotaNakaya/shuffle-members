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
	pid := ctx.Query("pid")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Tags []model.Tag
	if err := db.Where("project_id = ?", pid).Find(&Tags).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "tag/index.html", gin.H{"Tags": Tags, "PID": pid, "Project": Project})
}

// New はタグの新規作成画面に遷移します
func (t *TagController) New(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "tag/new.html", gin.H{"PID": pid, "Project": Project})
}

// Create はタグの作成を行います
func (t *TagController) Create(ctx *gin.Context) {
	db := db.GetDB()

	// TODO: バリデーション
	n := ctx.PostForm("name")
	pid := ctx.PostForm("pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	tag := model.Tag{
		Name:      n,
		ProjectID: pidInt,
	}
	if err := db.Create(&tag).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.Redirect(302, "/tag/index?pid="+pid)
}

// Edit はタグの編集画面に遷移します
func (t *TagController) Edit(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Tag model.Tag
	if err := db.Where("id = ?", id).Find(&Tag).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "tag/edit.html", gin.H{"PID": pid, "Tag": Tag, "Project": Project})
}

// Update はタグ情報の更新を行います
func (t *TagController) Update(ctx *gin.Context) {
	db := db.GetDB()
	// TODO: バリデーション

	id := ctx.Param("id")
	tag := model.Tag{}
	if err := db.First(&tag, id).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	n := ctx.PostForm("name")
	pid := ctx.PostForm("pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	db.Model(&tag).Updates(model.Tag{
		ProjectID: pidInt,
		Name:      n,
	})

	ctx.Redirect(302, "/tag/index?pid="+pid)
}

// Delete はタグの削除を行います
func (t *TagController) Delete(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var tag model.Tag
	if err := db.Delete(&tag, id).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.Redirect(302, "/tag/index?pid="+pid)
}
