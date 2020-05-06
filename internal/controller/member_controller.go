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
func (m *MemberController) Index(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Members []model.Member
	if err := db.Where("project_id = ?", pid).Preload("Tags").Find(&Members).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "member/index.html", gin.H{"Members": Members, "PID": pid, "Project": Project})
}

// New はメンバーの新規作成画面に遷移します
func (m *MemberController) New(ctx *gin.Context) {
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

	ctx.HTML(200, "member/new.html", gin.H{"PID": pid, "Tags": Tags, "Project": Project})
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
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// TODO: ハイパーやっつけ
	var t []model.Tag
	for i := 1; i <= 3; i++ {
		if tid := ctx.PostForm("tag" + strconv.Itoa(i)); tid != "" {
			tid, err := strconv.Atoi(tid)
			if err != nil {
				fmt.Println(err)
				ctx.HTML(500, "500.html", gin.H{"Error": err})
				return
			}
			t = append(t, model.Tag{ID: tid})
		}
	}

	member := model.Member{
		ProjectID: pidInt,
		Name:      n,
		Email:     e,
		Tags:      t,
	}
	if err := db.Create(&member).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.Redirect(302, "/member/index?pid="+pid)
}

// Edit はメンバーの編集画面に遷移します
func (m *MemberController) Edit(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Member model.Member
	if err := db.Where("id = ?", id).Preload("Tags").Find(&Member).Error; err != nil {
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

	ctx.HTML(200, "member/edit.html", gin.H{"PID": pid, "Member": Member, "Tags": Tags, "Project": Project})
}

// Update はメンバー情報の更新を行います
func (m *MemberController) Update(ctx *gin.Context) {
	db := db.GetDB()
	// TODO: バリデーション

	id := ctx.Param("id")
	member := model.Member{}
	if err := db.First(&member, id).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	n := ctx.PostForm("name")
	e := ctx.PostForm("email")
	pid := ctx.PostForm("pid")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// TODO: ハイパーやっつけ
	var t []model.Tag
	for i := 1; i <= 3; i++ {
		if tid := ctx.PostForm("tag" + strconv.Itoa(i)); tid != "" {
			tid, err := strconv.Atoi(tid)
			if err != nil {
				fmt.Println(err)
				ctx.HTML(500, "500.html", gin.H{"Error": err})
				return
			}
			t = append(t, model.Tag{ID: tid})
		}
	}

	db.Model(&member).Updates(model.Member{
		ProjectID: pidInt,
		Name:      n,
		Email:     e,
	}).Association("Tags").Replace(t)

	ctx.Redirect(302, "/member/index?pid="+pid)
}

// Delete はメンバーの削除を行います
func (m *MemberController) Delete(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var member model.Member
	if err := db.Delete(&member, id).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.Redirect(302, "/member/index?pid="+pid)
}
