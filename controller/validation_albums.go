package controller

import (
	"errors"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type AlbumsValidation struct{}

// アルバム情報のバリデーションを行う
func (v *AlbumsValidation) ValidateAlbum(album *model.Album) error {

	// パラメーターが不足している場合はエラー
	if album.ID == 0 {
		return errors.New("ID is required")
	}
	if album.Title == "" {
		return errors.New("album Title is required")
	}
	if album.SingerID == 0 {
		return errors.New("SingerID is required")
	}

	return nil
}
