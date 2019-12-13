package model

// Project はプロジェクトの情報を表します
type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
