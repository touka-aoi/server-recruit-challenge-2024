package repository

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

// Album用のCRUDのインターフェース
type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.Album, error)
	Get(ctx context.Context, id model.AlbumID) (*model.Album, error)
	Add(ctx context.Context, Album *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
}
