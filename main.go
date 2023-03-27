// エントリポイント
package main

// インポート
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pulse227/server-recruit-challenge-sample/api"
)

func main() {
	// interruptシグナルを受信したときに、コンテキストにキャンセルを通知する
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Routerの作成
	r := api.NewRouter()

	// HTTPサーバーの作成
	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}

	// ゴルーチンの作成
	// Graceful Shutdown
	go func() {
		// コンテキストのキャンセル通知を待機
		<-ctx.Done()
		// タイムアウト用のコンテキスト作成
		// 5s待機してからシャットダウンを行う
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// サーバーシャットダウン
		server.Shutdown(ctx)
	}()
	log.Println("server start running at :8888")
	// サーバーの起動
	log.Fatal(server.ListenAndServe())
}
