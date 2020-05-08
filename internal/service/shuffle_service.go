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
	var members []*model.Member
	if err := db.Where("project_id = ?", pid).Preload("ShuffleLogDetail").Preload("Tags").Find(&members).Error; err != nil {
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

	result := groupingMember(targetMemberList, gcount, mcount)

	d := []model.ShuffleLogDetail{}
	for k, v := range result {
		for _, v := range v {
			detail := model.ShuffleLogDetail{
				Group:    k,
				MemberID: v.ID,
			}
			d = append(d, detail)
		}
	}

	return d, nil
}

func makeTargetMemberList(members []*model.Member, length int) []*model.Member {
	// シャッフルログからメンバーとシャッフル回数のマップを作る
	memberCountMap := map[*model.Member]int{}
	for _, m := range members {
		memberCountMap[m] = len(m.ShuffleLogDetail)
	}

	// 回数とメンバーのマップの配列を作る
	countMembersMap := map[int][]*model.Member{}
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
	var targetMemberList []*model.Member
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
func shuffleSlice(data []*model.Member) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

// TODO: 全体的に汚い
func groupingMember(members []*model.Member, gcount, mcount int) map[int][]model.Member {
	result := map[int][]model.Member{}

	for i := 1; i <= gcount; i++ {
		fmt.Println("gcountの")
		fmt.Println(i)
		var m []model.Member
		var tags []model.Tag

		for i := 1; len(m) < mcount; i++ {
			fmt.Println("mcountの")
			fmt.Println(i)
			var target *model.Member
			// 対象を取り出す(popするけどcontinueの可能性もあるので、membersを更新しない)
			target, _ = pop(members)

			// タグ重複がないかチェック
			targetTags := target.Tags
			var b bool
			for _, v := range targetTags {
				b = contains(tags, v)
				if b == true {
					break
				}
			}

			// ループしきってないかつタグ重複の時はcontinue
			if i <= len(members) && b == true {
				continue
			}

			// membersを更新するためのpop
			target, members = pop(members)
			for _, v := range targetTags {
				tags = append(tags, v)
			}
			m = append(m, *target)
		}

		result[i] = m
	}
	return result
}

func pop(slice []*model.Member) (*model.Member, []*model.Member) {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans, slice
}

func contains(s []model.Tag, e model.Tag) bool {
	for _, v := range s {
		if e.ID == v.ID {
			return true
		}
	}
	return false
}
