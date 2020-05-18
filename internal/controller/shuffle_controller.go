package controller

import (
	"log"
	"sort"
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

	pid, err := strconv.Atoi(ctx.PostForm("pid"))
	gcount, err := strconv.Atoi(ctx.PostForm("gcount"))
	mcount, err := strconv.Atoi(ctx.PostForm("mcount"))
	// 雑
	if err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	service := service.ShuffleService{}
	ShuffleLogDetail, err := service.Shuffle(pid, gcount, mcount)
	if err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// ログインサート
	log := model.ShuffleLogHead{
		ProjectID:        pid,
		GroupCount:       gcount,
		MemberCount:      mcount,
		ShuffleLogDetail: ShuffleLogDetail,
	}
	if err := db.Create(&log).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// 雑にリダイレクト
	ctx.Redirect(302, "/shuffle/index?pid="+ctx.PostForm("pid"))
}

// Index はシャッフル結果の一覧を取得します
func (s *ShuffleController) Index(ctx *gin.Context) {
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
	if err := db.Debug().Where("project_id = ?", pid).Preload("ShuffleLogDetail").Find(&Logs).Error; err != nil {
		log.Print(err)
		ctx.HTML(500, "500.html", gin.H{"Error": err})
		return
	}

	// シャッフルログを表示用に整形する
	// 表示内容用のmap
	Result := map[string]map[int][]string{}
	// 表示順を保持した配列
	ResultKeys := []string{}
	for _, head := range Logs {
		m := map[int][]string{}
		for _, detail := range head.ShuffleLogDetail {
			v, _ := m[detail.Group]
			m[detail.Group] = append(v, memberNameByID(member, detail.MemberID))
		}
		c := head.CreatedAt.Format("2006-01-02T15:04")
		Result[c] = m
		ResultKeys = append(ResultKeys, c)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(ResultKeys)))

	ctx.HTML(200, "shuffle/index.html", gin.H{"Result": Result, "ResultKeys": ResultKeys, "Project": Project})
}

// メンバーIDに該当するメンバー名称を返す、存在しない場合はIDを返す
// TODO: これアソシエーションうまく使ってもっとスマートにやりたい
func memberNameByID(m []model.Member, id int) string {
	res := strconv.Itoa(id)
	for _, v := range m {
		if v.ID == id {
			res = v.Name
			break
		}
	}
	return res
}
