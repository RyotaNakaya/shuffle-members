package model

import "time"

// Project はプロジェクトの情報を表します
type Project struct {
	ID          int    `gorm:"primary_key; AUTO_INCREMENT"`
	Name        string `gorm:"unique; not null; size:50;"`
	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
