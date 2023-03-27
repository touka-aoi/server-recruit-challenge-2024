package model

// 歌手スキーマの定義

type SingerID int

// Singer構造体 Json時にはlowwerに
type Singer struct {
	ID   SingerID `json:"id"`
	Name string   `json:"name"`
}
