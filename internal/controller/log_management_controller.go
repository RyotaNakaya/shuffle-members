package controller

import (
	"log"
	"strconv"

	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/gin-gonic/gin"

	"github.com/RyotaNakaya/shuffle-members/db"
)

// LogManagementController はシャッフルログにの操作を行います
type LogManagementController struct {
}

// Index はシャッフルログの一覧を取得します
func (lm *LogManagementController) Index(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}
	var member []model.Member
	if err := db.Where("project_id = ?", pid).Find(&member).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var Logs []model.ShuffleLogHead
	if err := db.Debug().Order("id desc").Where("project_id = ?", pid).Preload("ShuffleLogDetail").Find(&Logs).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "log_management/index.html", gin.H{"Logs": Logs, "Project": Project})
}

// Edit はシャッフルログ詳細の編集画面に遷移します
func (lm *LogManagementController) Edit(ctx *gin.Context) {
	db := db.GetDB()
	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var Project model.Project
	if err := db.First(&Project, pid).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	var LogDetail model.ShuffleLogDetail
	if err := db.Where("id = ?", id).Find(&LogDetail).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.HTML(200, "log_management/edit.html", gin.H{"PID": pid, "LogDetail": LogDetail, "Project": Project})
}

// Update はシャッフルログ詳細の更新を行います
func (lm *LogManagementController) Update(ctx *gin.Context) {
	db := db.GetDB()
	// TODO: バリデーション

	id := ctx.Param("id")
	var logDetail model.ShuffleLogDetail
	if err := db.First(&logDetail, id).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	pid := ctx.PostForm("pid")
	g, err := strconv.Atoi(ctx.PostForm("group"))
	m, err := strconv.Atoi(ctx.PostForm("member_id"))
	if err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	db.Model(&logDetail).Updates(model.ShuffleLogDetail{
		Group:    g,
		MemberID: m,
	})

	ctx.Redirect(302, "/log_management/index?pid="+pid)
}

// Delete はシャッフルログ詳細の削除を行います
func (lm *LogManagementController) Delete(ctx *gin.Context) {
	db := db.GetDB()

	pid := ctx.Query("pid")
	id := ctx.Param("id")

	var detail model.ShuffleLogDetail
	if err := db.Delete(&detail, id).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	ctx.Redirect(302, "/log_management/index?pid="+pid)
}
