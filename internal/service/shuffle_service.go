package service

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/RyotaNakaya/shuffle-members/db"
	"github.com/RyotaNakaya/shuffle-members/internal/model"
)

// ShuffleService はシャッフル用の処理を行います
type ShuffleService struct {
}

// Shuffle はメンバーをシャッフルして結果を返します
func (s *ShuffleService) Shuffle(pid, gcount, mcount int) ([]model.ShuffleLogDetail, error) {
	db := db.GetDB()

	// そのプロジェクトに紐付くMemberの一覧を取得する
	var members []model.Member
	if err := db.Where("project_id = ?", pid).Preload("ShuffleLogDetail").Find(&members).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if gcount*mcount > len(members) {
		return nil, errors.New("入力された対象人数が多すぎます")
	}

	// 今回のシャッフル対象者リストを作る
	targetMemberList := makeTargetMemberList(members, gcount*mcount)
	if len(targetMemberList) != gcount*mcount {
		fmt.Printf("シャッフル対象者リストの数がおかしい expect: %d, return: %d", gcount*mcount, len(targetMemberList))
		return nil, errors.New("シャッフル対象者リストの抽出で問題が発生しました")
	}

	// TODO: タグを考慮してシャッフルする
	// 一人目のタグの人を除外
	// 二人目を決める
	// 以下ループ
	// 不足がある場合、以下をループする
	// 既にアサイン済みのIDは除いて最小回数の人たちを取得
	// タグを考慮してシャッフルしてappendする

	d := []model.ShuffleLogDetail{}
	group := 1
	counter := 0
	for k := range targetMemberList {
		detail := model.ShuffleLogDetail{
			Group:    group,
			MemberID: targetMemberList[k],
		}
		d = append(d, detail)

		counter++
		if counter == mcount {
			group++
			counter = 0
		}
	}

	return d, nil
}

func makeTargetMemberList(members []model.Member, length int) []int {
	// シャッフルログからメンバーIDとシャッフル回数のマップを作る
	memberCountMap := map[int]int{}
	for _, m := range members {
		memberCountMap[m.ID] = len(m.ShuffleLogDetail)
	}

	// 回数とメンバーIDのマップの配列を作る
	countMembersMap := map[int][]int{}
	for user, count := range memberCountMap {
		v, _ := countMembersMap[count]
		countMembersMap[count] = append(v, user)
	}

	// 当選回数でソートされた配列を作る
	var keys []int
	for key := range countMembersMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// 回数の少ない人から抽出していって、今回の対象者リストを作る
	var targetMemberList []int
	for _, key := range keys {
		if len(targetMemberList) >= length {
			break
		}

		members := countMembersMap[key]
		shuffleSlice(members)
		for _, m := range members {
			if len(targetMemberList) >= length {
				break
			}
			targetMemberList = append(targetMemberList, m)
		}
	}

	return targetMemberList
}

// スライスをシャッフルして返す
func shuffleSlice(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
