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

	var Members []model.Member
	if err := db.Where("project_id = ?", pid).Find(&Members).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "member/index.html", gin.H{"Members": Members, "PID": pid})
}

// New はメンバーの新規作成画面に遷移します
func (m *MemberController) New(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	var Tags []model.Tag
	if err := db.Where("project_id = ?", pid).Find(&Tags).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "member/new.html", gin.H{"PID": pid, "Tags": Tags})
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
	err = createMemberTag(ctx, pidInt, member.ID)
	if err != nil {
		panic("fail createMemberTag")
	}

	ctx.Redirect(302, "/member/index?pid="+pid)
}

// Edit はメンバーの編集画面に遷移します
func (m *MemberController) Edit(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/member/index?pid="+pid)
	}

	var Member model.Member
	if err := db.Where("id = ?", id).Find(&Member).Error; err != nil {
		fmt.Println(err)
	}

	// TODO: MemberTagの取得
	// Member に []Tag フィールド持たせた方が楽か？

	var Tags []model.Tag
	if err := db.Where("project_id = ?", pid).Find(&Tags).Error; err != nil {
		fmt.Println(err)
	}

	ctx.HTML(200, "member/edit.html", gin.H{"PID": pid, "Member": Member, "Tags": Tags})
}

// Delete はメンバーの削除を行います
func (m *MemberController) Delete(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/member/index?pid="+pid)
	}

	var member model.Member
	if err := db.Delete(&member, id).Error; err != nil {
		// TODO: エラーハンドリング
		fmt.Println(err)
		ctx.Redirect(302, "/member/index?pid="+pid)
	}

	ctx.Redirect(302, "/member/index?pid="+pid)
}

func createMemberTag(ctx *gin.Context, pid int, mid int) error {
	// TODO: ハイパーやっつけ
	db := db.GetDB()
	for i := 1; i <= 3; i++ {
		if tid := ctx.PostForm("tag" + strconv.Itoa(i)); tid != "" {
			tid, err := strconv.Atoi(tid)
			weight, err := strconv.Atoi(ctx.PostForm("weight" + strconv.Itoa(i)))
			if err != nil {
				// TODO: エラーハンドリング
				panic("convert fail")
			}
			mt := model.MemberTag{
				ProjectID: pid,
				MemberID:  mid,
				TagID:     tid,
				Weight:    weight,
			}
			if err := db.Create(&mt).Error; err != nil {
				// TODO: エラーハンドリング
				fmt.Println(err)
				return (err)
			}
		}
	}

	return nil
}
