package main_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pulse227/server-recruit-challenge-sample/api"
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/stretchr/testify/assert"
)

// 4-1 指定したIDのアルバムを取得するAPI
func TestAlbumSingerGet(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
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

		expected := model.AlbumSinger{ID: 1, Title: "Alice's 1st Album", Singer: model.Singer{ID: 1, Name: "Alice"}}

		var album model.AlbumSinger
		// レスポンスのボディを確認
		if err := json.NewDecoder(rr.Body).Decode(&album); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expected, album)
	})
}

// 4-2 アルバムの一覧を取得するAPI
func TestAlbumSingerGetAll(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		// ルーターを作成
		r := api.NewRouter()
		// HTTPリクエストを作成
		req, err := http.NewRequest("GET", "/albums", nil)
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

		expected := []*model.AlbumSinger{
			{ID: 1, Title: "Alice's 1st Album", Singer: model.Singer{ID: 1, Name: "Alice"}},
			{ID: 2, Title: "Alice's 2nd Album", Singer: model.Singer{ID: 1, Name: "Alice"}},
			{ID: 3, Title: "Bella's 1st Album", Singer: model.Singer{ID: 2, Name: "Bella"}},
		}

		var albums []*model.AlbumSinger
		// レスポンスのボディを確認
		if err := json.NewDecoder(rr.Body).Decode(&albums); err != nil {
			t.Fatal(err)
		}

		assert.ElementsMatch(t, expected, albums)
	})
}
