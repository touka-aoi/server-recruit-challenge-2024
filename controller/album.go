package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

type albumController struct {
	service service.AlbumService
}

// コンストラクタ
func NewAlbumController(service service.AlbumService) *albumController {
	return &albumController{service: service}
}

// GET /albums のハンドラ
func (c *albumController) GetAlbumListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumListService(r.Context())
	if err != nil {
		// エラーの場合サーバーエラーを出して終了
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンスの作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// JSONレスポンス
	json.NewEncoder(w).Encode(albums)
}

// GET /albums/{id} のハンドラ
func (c *albumController) GetAlbumDetailHandler(w http.ResponseWriter, r *http.Request) {
	// パスパラメータの取得
	albumID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// エラー処理
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}
	// 特定のアルバムの呼び出し
	album, err := c.service.GetAlbumService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンス作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// POST /albums のハンドラ
func (c *albumController) PostAlbumHandler(w http.ResponseWriter, r *http.Request) {
	var album *model.Album

	// リクエストのパース
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		// パース時のエラー処理
		err = fmt.Errorf("invalid request body: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	// リクエストのバリデーション
	validation := &AlbumsValidation{}

	if err := validation.ValidateAlbum(album); err != nil {
		errorHandler(w, r, 400, err.Error())
		return
	}

	// アルバムの作成
	if err := c.service.PostAlbumService(r.Context(), album); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	// レスポンス作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}

// DELETE /albums/{id} のハンドラ
func (c *albumController) DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	// パスパラメータの取得
	albumID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// エラー処理
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	// アルバムの削除
	if err := c.service.DeleteAlbumService(r.Context(), model.AlbumID(albumID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	// レスポンス作成
	w.WriteHeader(204)
}
