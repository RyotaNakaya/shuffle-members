package service

import (
	"fmt"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/model"
)

// ShuffleService はシャッフル用の処理を行います
type ShuffleService struct {
}

// Shuffle はメンバーをシャッフルして結果を返します
func (s *ShuffleService) Shuffle(pid, gcount, mcount int) ([]model.Member, error) {
	db := db.GetDB()

	// まずはログを取得
	var log []model.ShuffleLog
	if err := db.Where("project_id = ?", pid).Find(&log).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 回数を数えて最小回数の人たちを取得
	// タグを考慮してシャッフルする
	// 一人目のタグの人を除外
	// 二人目を決める
	// 以下ループ
	// 不足がある場合、以下をループする
	// 既にアサイン済みのIDは除いて最小回数の人たちを取得
	// タグを考慮してシャッフルしてappendする

	m := []model.Member{}
	return m, nil
}
