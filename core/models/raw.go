package models

// 入库前的结构
type ArtworkRaw struct {
	ID          uint
	Title       string
	Author      string
	Description string
	Source      SourceName
	SourceURL   string
	Tags        []string
	R18         bool
	Pictures    []*PictureRaw
}

type PictureRaw struct {
	DirectURL  string
	Width      uint
	Height     uint
	Hash       string
	Format     string
	Binary     []byte
	BlurScore  float64
	FilePath   string
	Downloaded bool
}
