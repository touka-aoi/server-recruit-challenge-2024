package api_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pulse227/server-recruit-challenge-sample/api"
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/stretchr/testify/assert"
)

// GET /Albums のテスト
func TestAlbumGetAll(t *testing.T) {
	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", "/albums", nil)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr := httptest.NewRecorder()
	// ルーターを作成
	r := api.NewRouter()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	log.Print("req: " + rr.Body.String())

	expected := []*model.Album{
		{ID: 1, Title: "Alice's 1st Album", SingerID: 1},
		{ID: 2, Title: "Alice's 2nd Album", SingerID: 1},
		{ID: 3, Title: "Bella's 1st Album", SingerID: 2},
	}

	var albums []*model.Album
	// レスポンスのボディを確認
	if err := json.NewDecoder(rr.Body).Decode(&albums); err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, expected, albums)

}

// GET /albums{id} のテスト
func TestAlbumGet(t *testing.T) {
	// ルーターを作成
	r := api.NewRouter()
	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", "/albums/1", nil)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr := httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	log.Print("req: " + rr.Body.String())

	expected := &model.Album{ID: 1, Title: "Alice's 1st Album", SingerID: 1}

	var albums *model.Album
	// レスポンスのボディを確認
	if err := json.NewDecoder(rr.Body).Decode(&albums); err != nil {
		t.Fatal(err)
	}

	// レスポンスのボディを確認
	assert.Equal(t, expected, albums)

	// 存在しないIDの場合 500BadRequestが来る
	// Todo: 404NotFoundが来るべき
	req, err = http.NewRequest("GET", "/albums/1000", nil)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認 { 500 internal server error が返ってくる }
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// 負数の場合 404 not found が返る
	req, err = http.NewRequest("GET", "/albums-1", nil)

	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認 { 404 not foundが返ってくる }
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

}

// POST /albums のテスト
// POSTリクエストはリクエストデータを登録することのみを行う
// バリデーションなどはServiceレベルで行う
func TestAlbumPost(t *testing.T) {
	// ルーターを作成
	r := api.NewRouter()

	// 重複しないIDかつ、存在するSingerIDの場合 //

	// リクエストを作成
	parm := `{"id": 4, "title":"Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body := []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody := bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr := httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// レスポンスのボディを確認
	log.Print("TEST1 正常POST: " + rr.Body.String())

	expected := &model.Album{ID: 4, Title: "Alice's 3rd Album", SingerID: 1}

	var albums *model.Album
	// レスポンスのボディを確認
	if err := json.NewDecoder(rr.Body).Decode(&albums); err != nil {
		t.Fatal(err)
	}

	// レスポンスのボディを確認
	assert.Equal(t, expected, albums)

	// IDが重複した場合はそのまま上書きをする //

	// リクエストを作成
	parm = `{"id": 1, "title":"Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// レスポンスのボディを確認
	log.Print("TEST2 上書きPOST: " + rr.Body.String())

	expected = &model.Album{ID: 1, Title: "Alice's 3rd Album", SingerID: 1}

	// レスポンスのボディを確認
	if err := json.NewDecoder(rr.Body).Decode(&albums); err != nil {
		t.Fatal(err)
	}

	// IDが文字列の場合はエラーを返す //

	// リクエストを作成
	parm = `{"id": "1", "title": "Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST3 タイプエラー Code: ", rr.Code, " ", rr.Body.String())

	// IDが0の場合はエラーを返す //

	// リクエストを作成
	parm = `{"id": 0, "title": "Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST4 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// titleが空の場合はエラーを返す

	// リクエストを作成
	parm = `{"id": 0, "title": "","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST5 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// titleが文字列以外の場合はエラーを返す
	// リクエストを作成
	parm = `{"id": 1, "title": 11111,"singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST6 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// IDがnullの場合エラーを返す
	// リクエストを作成
	parm = `{"id": null, "title":"Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST7 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// リクエストボディが不足している場合エラーを返す
	// リクエストを作成
	parm = `{}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)
	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST8 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// リクエストを作成
	parm = `{"id": 1, "title":"Alice's 3rd Album"}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST9 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// リクエストを作成
	parm = `{"title":"Alice's 3rd Album","singer_id":1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST10 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// リクエストを作成
	parm = `{"id": 1}`
	// JSON文字列をバイトスライスに変換
	body = []byte(parm)
	// io.Reader型のオブジェクトを作成
	requestBody = bytes.NewBuffer(body)

	// HTTPリクエストを作成
	req, err = http.NewRequest("POST", "/albums", requestBody)
	// 作成に失敗した場合
	if err != nil {
		t.Fatal(err)
	}
	// レスポンスを用意
	rr = httptest.NewRecorder()
	// ルーターにリクエストを送信
	r.ServeHTTP(rr, req)

	// レスポンスのステータスコードを確認
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	log.Print("TEST11 REQUIRE Code: ", rr.Code, " ", rr.Body.String())

	// t.Fatal("TestAlbumsPost: ここまで到達したらエラー")

}

// アルバムを削除する
// Test: /albumsN delete Test

// 存在する

// 存在しない
