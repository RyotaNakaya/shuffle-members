package controller

import (
	"fmt"
	"strconv"

	"github.com/RyotaNakaya/shuffle-members/internal/model"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/service"
	"github.com/gin-gonic/gin"
)

// ShuffleController はシャッフルに関する処理を行います
type ShuffleController struct {
}

// Shuffle はメンバーをシャッフルして結果を返します
func (s *ShuffleController) Shuffle(ctx *gin.Context) {
	db := db.GetDB()

	if ctx.PostForm("gcount") == "" {
		ctx.HTML(500, "500.html", gin.H{"Error": "グループ数は必須です"})
		return
	} else if ctx.PostForm("mcount") == "" {
		ctx.HTML(500, "500.html", gin.H{"Error": "人数は必須です"})
		return
	}

	pid, err := strconv.Atoi(ctx.PostForm("id"))
	gcount, err := strconv.Atoi(ctx.PostForm("gcount"))
	mcount, err := strconv.Atoi(ctx.PostForm("mcount"))
	// 雑
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	service := service.ShuffleService{}
	Members, err := service.Shuffle(pid, gcount, mcount)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// ログインサート
	log := model.ShuffleLog{
		ProjectID:   pid,
		GroupCount:  gcount,
		MemberCount: mcount,
		Members:     Members,
	}
	if err := db.Create(&log).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// 雑にリダイレクト
	ctx.Redirect(302, "/shuffle/index")
}

// Index はシャッフル結果の一覧を取得します
func (s *ShuffleController) Index(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Param("id")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Logs []model.ShuffleLog
	if err := db.Where("project_id = ?", pid).Find(&Logs).Error; err != nil {
		fmt.Println(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "shuffle/index.html", gin.H{"Logs": Logs, "Project": Project})
}
