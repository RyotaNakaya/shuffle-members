package service

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/RyotaNakaya/shuffle-members/db"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	err      error
	fixtures *testfixtures.Context
)

func TestShuffle(t *testing.T) {
	t.Run("too many input count", func(t *testing.T) {
		prepareTestDatabase()

		projectID, groupCount, memberCount := 1, 10, 10
		s := ShuffleService{}
		_, err := s.Shuffle(projectID, groupCount, memberCount)
		if err == nil {
			t.Error("error is expeted")
		}
	})
	t.Run("happy path", func(t *testing.T) {
		prepareTestDatabase()

		projectID, groupCount, memberCount := 1, 2, 2
		s := ShuffleService{}
		res, err := s.Shuffle(projectID, groupCount, memberCount)
		if err != nil {
			t.Error(err)
		}
		// レスポンス数チェック
		if len(res) != groupCount*memberCount {
			t.Errorf("response length wrong, expected: %d, got: %d", groupCount*memberCount, len(res))
		}
		// 同じタグを持ってる人が重らないかチェック
		var g5, g6 int
		for _, v := range res {
			if v.MemberID == 5 {
				g5 = v.Group
			}
			if v.MemberID == 6 {
				g6 = v.Group
			}
		}
		if g5 == g6 {
			t.Errorf("tag logic is broken, res: %+v", res)
		}
		// シャッフルしきって跨いでも大丈夫かチェック
		projectID, groupCount, memberCount = 1, 2, 3
		if _, err = s.Shuffle(projectID, groupCount, memberCount); err != nil {
			t.Error(err)
		}
		projectID, groupCount, memberCount = 1, 2, 2
		if _, err = s.Shuffle(projectID, groupCount, memberCount); err != nil {
			t.Error(err)
		}
	})
}

func init() {
	// 環境変数のセット
	os.Setenv("MYSQL_DB", "shuffle_members_test")
	os.Setenv("MYSQL_DB_HOST", "")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "")

	// DB セットアップ
	db.Init()
	// defer db.Close()

	// テストデータの準備
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("MYSQL_DB_HOST")
	DBNAME := os.Getenv("MYSQL_DB")
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"
	cnnect := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	prepareFixture(cnnect)
	prepareTestDatabase()
}

func prepareFixture(dbstr string) {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		log.Fatal(err)
	}

	fixtures, err = testfixtures.NewFolder(db, &testfixtures.MySQL{}, "../../testdata/fixture")
	if err != nil {
		log.Fatal(err)
	}
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}

func Init() {
	// 環境変数のセット
	os.Setenv("MYSQL_DB", "shuffle_members_test")
	os.Setenv("MYSQL_DB_HOST", "")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "")

	// DB セットアップ
	db.Init()
	// defer db.Close()

	// テストデータの準備
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("MYSQL_DB_HOST")
	DBNAME := os.Getenv("MYSQL_DB")
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"
	cnnect := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	prepareFixture(cnnect)
	prepareTestDatabase()
}
