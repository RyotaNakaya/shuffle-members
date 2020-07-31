package model

import "time"

// ShuffleLogHead はシャッフル情報のログを表します
type ShuffleLogHead struct {
	ID               int `gorm:"primary_key; AUTO_INCREMENT"`
	ProjectID        int `sql:"index"`
	GroupCount       int
	MemberCount      int
	ShuffleLogDetail []ShuffleLogDetail

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ShuffleLogDetail はシャッフルのメンバー情報を表します
type ShuffleLogDetail struct {
	ID               int `gorm:"primary_key; AUTO_INCREMENT"`
	ShuffleLogHeadID int `sql:"index"`
	Group            int
	MemberID         int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
