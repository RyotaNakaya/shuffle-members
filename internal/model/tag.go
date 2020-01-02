package model

import "time"

// Tag はタグの情報を表します
type Tag struct {
	ID        int    `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	ProjectID int    `sql:"index"`
	Name      string `gorm:"unique; not null; size:50;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// FOREIGN KEY (project_id)
//         references project (id)
//         on delete restrict
//         on update restrict
