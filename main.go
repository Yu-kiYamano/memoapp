package main

import (
	"memoapp/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

func main() {

	startServer() //サーバースタート
}

func startServer() *echo.Echo {

	e := echo.New() //echoインスタンス作成
	// dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	// db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	// if err != nil {
	// 	e.Logger.Errorf("データベース接続に失敗しました。: %v\n", err)
	// }
	// if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
	// 	e.Logger.Errorf("確認できません: %v\n", err)
	// }
	// log.Println("データベースに接続しました")

	// defer db.Close() //startServer関数が終了する際に実行されるようにする為、deferを記述

	hdlr := handler.ProvideMemohandler()
	//middlewareを登録
	e.Use(
		middleware.Recover(), //パニックから回復させるためのmiddleware
		middleware.Logger(),  //各HTTP リクエストに関する情報をログに記録するためのmiddleware
		middleware.Gzip(),    //gzip圧縮スキームを使用してHTTPレスポンスを圧縮するためのmiddleware
		hdlr.CheckCache(),
	)

	e.POST("/", hdlr.MemoCreate)
	e.GET("/", hdlr.MemoIndex)
	e.DELETE("/:id", hdlr.MemoDelete)
	e.Logger.Fatal(e.Start(":8080"))
	return e

}
