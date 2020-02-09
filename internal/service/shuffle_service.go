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
func (s *ShuffleService) Shuffle(pid, gcount, mcount int) ([]model.ShuffleLogDetail, error) {
	db := db.GetDB()

	// そのプロジェクトに紐付くシャッフルログのヘッダとディテールを取得する
	var log []model.ShuffleLogHead
	if err := db.Debug().Where("project_id = ?", pid).Preload("ShuffleLogDetail").Find(&log).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(log)

	// 回数を数えて最小回数の人たちを取得
	// detail のメンバーIDのSUMを取得し、一番少ないメンバーのリストを取得
	// タグを考慮してシャッフルする
	// 一人目のタグの人を除外
	// 二人目を決める
	// 以下ループ
	// 不足がある場合、以下をループする
	// 既にアサイン済みのIDは除いて最小回数の人たちを取得
	// タグを考慮してシャッフルしてappendする

	// m := []model.Member{}
	m := []model.ShuffleLogDetail{}
	m1 := model.ShuffleLogDetail{MemberID: 1}
	m2 := model.ShuffleLogDetail{MemberID: 2}
	m = append(m, m1)
	m = append(m, m2)
	return m, nil
}
