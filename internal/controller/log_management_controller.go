package controller

import (
	"log"

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
