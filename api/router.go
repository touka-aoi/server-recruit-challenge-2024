package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/memorydb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter() *mux.Router {
	// 歌手情報
	// 歌手DBの作成
	singerRepo := memorydb.NewSingerRepository()
	// 歌手サービスの作成
	singerService := service.NewSingerService(singerRepo)
	// 歌手コントローラの作成
	singerController := controller.NewSingerController(singerService)

	// アルバム情報
	// アルバムDBの作成
	albumRepo := memorydb.NewAlbumRepository()
	// アルバムサービスの作成
	albumService := service.NewAlbumService(albumRepo)
	// アルバムコントローラの作成 (課題3の場合はこっち)
	// albumController := controller.NewAlbumController(albumService)

	// アルバム + 歌手情報
	albumSingerService := service.NewAlbumSingerService(albumService, singerService)
	// アルバムコントローラの作成
	// (課題4の場合はこっち)
	albumController := controller.NewAlbumSingerController(albumSingerService)

	// ルータの作成
	r := mux.NewRouter()

	// ルーター設定
	// 歌手
	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet) // GET /singers
	r.HandleFunc("/singers/{id:[1-9][0-9]*}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[1-9][0-9]*}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)
	// アルバム
	r.HandleFunc("/albums", albumController.GetAlbumListHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id:[1-9][0-9]*}", albumController.GetAlbumDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums", albumController.PostAlbumHandler).Methods(http.MethodPost)
	r.HandleFunc("/albums/{id:[1-9][0-9]*}", albumController.DeleteAlbumHandler).Methods(http.MethodDelete)

	// ミドルウェアの設定 (ログ出力)
	r.Use(middleware.LoggingMiddleware)

	return r
}
