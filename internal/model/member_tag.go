package model

import "time"

// MemberTag はメンバーとタグのマッピングを表します
type MemberTag struct {
	ID        int `gorm:"AUTO_INCREMENT" json:"id"`
	ProjectID int
	MemberID  int
	TagID     int
	Weight    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
