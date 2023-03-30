package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumSingerService interface {
	GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error)
	GetAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) (*model.AlbumSinger, error)
	PostAlbumSingerService(ctx context.Context, Album *model.Album) error
	DeleteAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) error
}

type albumSingerService struct {
	albumRepository  repository.AlbumRepository
	singerRepository repository.SingerRepository
}

var _ AlbumSingerService = (*albumSingerService)(nil)

func NewAlbumSingerService(albumRepository repository.AlbumRepository, singerRepository repository.SingerRepository) AlbumSingerService {
	return &albumSingerService{
		albumRepository:  albumRepository,
		singerRepository: singerRepository,
	}
}

func (s *albumSingerService) GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error) {
	// アルバムデータの取得
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// レスポンスデータの初期化
	albumsSinger := make([]*model.AlbumSinger, 0, len(albums))

	// 歌手データの取得
	for _, album := range albums {
		singer, err := s.singerRepository.Get(ctx, album.SingerID)
		if err != nil {
			return nil, err
		}

		// アルバムと歌手のデータを結合
		albumsSinger = append(albumsSinger, &model.AlbumSinger{
			ID:     album.ID,
			Title:  album.Title,
			Singer: *singer,
		})
	}

	return albumsSinger, nil
}

func (s *albumSingerService) GetAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) (*model.AlbumSinger, error) {
	album, err := s.albumRepository.Get(ctx, AlbumID)
	if err != nil {
		return nil, err
	}

	singer, err := s.singerRepository.Get(ctx, album.SingerID)
	if err != nil {
		return nil, err
	}

	return &model.AlbumSinger{
		ID:     album.ID,
		Title:  album.Title,
		Singer: *singer,
	}, nil
}

func (s *albumSingerService) PostAlbumSingerService(ctx context.Context, Album *model.Album) error {
	return nil
}

func (s *albumSingerService) DeleteAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) error {
	return nil
}
