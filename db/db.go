package db

import (
	"github.com/RyotaNakaya/shuffle-members/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use mysql in gorm
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open(getDBConfig())
	if err != nil {
		panic("failed to connect database")
	}

	// スキーマのマイグレーション
	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(model.Tag{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(model.Member{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(model.ShuffleLogHead{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(model.ShuffleLogDetail{}).AddForeignKey("shuffle_log_head_id", "shuffle_log_heads(id)", "RESTRICT", "RESTRICT")
}

// GetDB は gorm.DB インスタンスを返します
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func getDBConfig() (string, string) {
	// TODO: 設定ファイルから読み込む
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(docker.for.mac.localhost:3306)"
	DBNAME := "shuffle_members_development"
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}
