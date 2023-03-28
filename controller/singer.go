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

type singerController struct {
	service service.SingerService
}

// コンストラクタ
func NewSingerController(s service.SingerService) *singerController {
	return &singerController{service: s}
}

// GET /singers のハンドラー
func (c *singerController) GetSingerListHandler(w http.ResponseWriter, r *http.Request) {
	singers, err := c.service.GetSingerListService(r.Context())
	if err != nil {
		// エラーの場合サーバーエラーを出して終了
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンスの作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// JSONレスポンス
	json.NewEncoder(w).Encode(singers)
}

// GET /singers/{id} のハンドラー
func (c *singerController) GetSingerDetailHandler(w http.ResponseWriter, r *http.Request) {
	// パスパラメータの取得
	singerID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// エラー処理
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}
	// 特定の歌手の呼び出し
	singer, err := c.service.GetSingerService(r.Context(), model.SingerID(singerID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	// レスポンス作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// POST /singers のハンドラー
func (c *singerController) PostSingerHandler(w http.ResponseWriter, r *http.Request) {
	var singer *model.Singer
	// リクエストボディのパース・エラーチェック
	if err := json.NewDecoder(r.Body).Decode(&singer); err != nil {
		err = fmt.Errorf("invalid body param: %w", err) // 文字列エラーの作成
		errorHandler(w, r, 400, err.Error())            // 400 Bad Requestを返す
		return
	}

	// 歌手データの保存
	if err := c.service.PostSingerService(r.Context(), singer); err != nil {
		errorHandler(w, r, 500, err.Error()) // 500 Internal Server Errorを返す
		return
	}

	// レスポンスの作成
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// DELETE /singers/{id} のハンドラー
func (c *singerController) DeleteSingerHandler(w http.ResponseWriter, r *http.Request) {
	singerID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// エラーチェックを行い エラーの場合400 Bad Requestを返す
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	// 歌手データの削除
	if err := c.service.DeleteSingerService(r.Context(), model.SingerID(singerID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.WriteHeader(204) // 204 No Contentを返す
}
