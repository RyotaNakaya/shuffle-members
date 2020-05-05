package model

import (
	"time"
)

// Member はメンバーの情報を表します
type Member struct {
	ID               int    `gorm:"primary_key; AUTO_INCREMENT"`
	ProjectID        int    `sql:"index"`
	Name             string `gorm:"unique; not null; size:50;"`
	Email            string `gorm:"type:varchar(100);unique_index"`
	Tags             []Tag  `gorm:"many2many:member_tags;association_autoupdate:false"`
	ShuffleLogDetail []ShuffleLogDetail

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
