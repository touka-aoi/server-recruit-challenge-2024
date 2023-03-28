package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, AlbumID model.AlbumID) (*model.Album, error)
	PostAlbumService(ctx context.Context, Album *model.Album) error
	DeleteAlbumService(ctx context.Context, AlbumID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
}

// albumServiceがAlbumServiceを実装
var _ AlbumService = (*albumService)(nil)

// コンストラクタ
func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}

// GetAlbumListService
func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

// GetAlbumService
func (s *albumService) GetAlbumService(ctx context.Context, AlbumID model.AlbumID) (*model.Album, error) {
	album, err := s.albumRepository.Get(ctx, AlbumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

// PostAlbumService
func (s *albumService) PostAlbumService(ctx context.Context, Album *model.Album) error {
	if err := s.albumRepository.Add(ctx, Album); err != nil {
		return err
	}
	return nil
}

// DeleteAlbumService
func (s *albumService) DeleteAlbumService(ctx context.Context, AlbumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, AlbumID); err != nil {
		return err
	}
	return nil
}
