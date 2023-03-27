package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type singerRepository struct {
	sync.RWMutex
	singerMap map[model.SingerID]*model.Singer // キーが SingerID、値が model.Singer のマップ
}

var _ repository.SingerRepository = (*singerRepository)(nil)

// 初期データを作成する関数
func NewSingerRepository() *singerRepository {
	var initMap = map[model.SingerID]*model.Singer{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bella"},
		3: {ID: 3, Name: "Chris"},
		4: {ID: 4, Name: "Daisy"},
		5: {ID: 5, Name: "Ellen"},
	}

	return &singerRepository{
		singerMap: initMap,
	}
}

// すべての歌手を取得する
func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	// 書き込みの排他制御
	r.RLock()
	defer r.RUnlock()

	// サイズが歌手人数のスライスを作成する
	singers := make([]*model.Singer, 0, len(r.singerMap))
	for _, s := range r.singerMap {
		singers = append(singers, s)
	}
	return singers, nil
}

// IDから歌手を取得する
func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()

	singer, ok := r.singerMap[id]
	// インデックスが見つからない場合falseを返す
	if !ok {
		return nil, errors.New("not found")
	}
	return singer, nil
}

// 歌手を追加する
func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	r.Lock()
	r.singerMap[singer.ID] = singer
	r.Unlock()
	return nil
}

// 歌手を削除する
func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	r.Lock()
	delete(r.singerMap, id)
	r.Unlock()
	return nil
}
