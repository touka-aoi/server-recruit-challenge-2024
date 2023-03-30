package model

// アルバムスキーマの定義

type AlbumSinger struct {
	ID     AlbumID `json:"id"`
	Title  string  `json:"title"`
	Singer Singer  `json:"singer"`
}
