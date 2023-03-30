package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type AlbumSingerService interface {
	GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error)
	GetAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) (*model.AlbumSinger, error)
	PostAlbumSingerService(ctx context.Context, Album *model.Album) error
	DeleteAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) error
}

type albumSingerService struct {
	albumSvc  albumService
	singerSvc singerService
}

var _ AlbumSingerService = (*albumSingerService)(nil)

func NewAlbumSingerService(albumSvc *albumService, singerSvc *singerService) *albumSingerService {
	return &albumSingerService{
		albumSvc:  *albumSvc,
		singerSvc: *singerSvc,
	}
}

func (s *albumSingerService) GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error) {
	// アルバムデータの取得
	albums, err := s.albumSvc.GetAlbumListService(ctx)
	if err != nil {
		return nil, err
	}

	// レスポンスデータの初期化
	albumsSinger := make([]*model.AlbumSinger, 0, len(albums))

	// 歌手データの取得
	for _, album := range albums {
		singer, err := s.singerSvc.GetSingerService(ctx, album.SingerID)
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
	// アルバムデータの取得
	album, err := s.albumSvc.GetAlbumService(ctx, AlbumID)
	if err != nil {
		return nil, err
	}

	// 歌手データの取得
	singer, err := s.singerSvc.GetSingerService(ctx, album.SingerID)
	if err != nil {
		return nil, err
	}

	// アルバムと歌手のデータを結合
	return &model.AlbumSinger{
		ID:     album.ID,
		Title:  album.Title,
		Singer: *singer,
	}, nil

}

func (s *albumSingerService) PostAlbumSingerService(ctx context.Context, Album *model.Album) error {
	// アルバムデータの登録
	if err := s.albumSvc.PostAlbumService(ctx, Album); err != nil {
		return err
	}
	return nil
}

func (s *albumSingerService) DeleteAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) error {
	// アルバムデータの削除
	if err := s.albumSvc.DeleteAlbumService(ctx, AlbumID); err != nil {
		return err
	}
	return nil
}
