package repository

import (
	"log"
	"memoapp/model"

	"github.com/labstack/echo/v4"
)

var LocalCache = map[string][]*model.Memo{
	// "メモ一覧": []*model.Memo{
	// 	&model.Memo{
	// 		ID:   1,
	// 		Memo: "cacheMemo1",
	// 	},
	// 	&model.Memo{
	// 		ID:   2,
	// 		Memo: "cacheMemo2",
	// 	},
	// 	&model.Memo{
	// 		ID:   3,
	// 		Memo: "cacheMemo3",
	// 	},
	// },
}

type Cache struct {
}

var _ Database = Cache{}

func ProvieCache(c echo.Context) (Cache, error) {
	// dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	// db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	// if err != nil {
	// 	c.Logger().Errorf("データベース接続に失敗しました。: %v\n", err)
	// 	return Cache{}, err
	// }

	// if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
	// 	c.Logger().Errorf("確認できません: %v\n", err)
	// 	return Cache{}, err
	// }

	log.Println("ローカルキャッシュに接続しました")

	return Cache{}, nil
}

// func (cache Cache) Close() error {
// 	err := cache.db.Close()
// 	return err

// }

func (cache Cache) Set(c echo.Context, memo *model.Memo) error {
	// LocalCache["メモ一覧"] = memo
	getCache := LocalCache["メモ一覧"]
	getCache = append(getCache, memo)
	LocalCache["メモ一覧"] = getCache
	return nil
}

func (cache Cache) Get() ([]*model.Memo, error) {
	getCache := LocalCache["メモ一覧"]
	return getCache, nil
}

func (cach Cache) Judge() bool {
	_, ok := LocalCache["メモ一覧"]
	return ok
}

func (cache Cache) Delete(c echo.Context, id int) error {
	return nil
}
