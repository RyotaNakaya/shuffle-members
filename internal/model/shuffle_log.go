package model

import "time"

// ShuffleLog はシャッフル情報のログを表します
type ShuffleLog struct {
	ID          int `gorm:"primary_key; AUTO_INCREMENT"`
	ProjectID   int `sql:"index"`
	GroupCount  int
	MemberCount int
	Members     []Member

	CreatedAt time.Time
}
