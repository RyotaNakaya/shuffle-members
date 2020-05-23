package model

import (
	"time"
)

// Member はメンバーの情報を表します
type Member struct {
	ID               int    `gorm:"primary_key; AUTO_INCREMENT"`
	ProjectID        int    `sql:"index"`
	Name             string `gorm:"not null; size:50;"`
	Email            string `gorm:"type:varchar(100);"`
	Tags             []Tag  `gorm:"many2many:member_tags;association_autoupdate:false"`
	ShuffleLogDetail []ShuffleLogDetail

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
