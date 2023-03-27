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
	// 歌手DBの作成
	singerRepo := memorydb.NewSingerRepository()
	// 歌手サービスの作成
	singerService := service.NewSingerService(singerRepo)
	// 歌手コントローラの作成
	singerController := controller.NewSingerController(singerService)

	// ルータの作成
	r := mux.NewRouter()

	// ルーター設定
	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	r.Use(middleware.LoggingMiddleware)

	return r
}
