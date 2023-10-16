package models

import "gorm.io/gorm"

type Picture struct {
	gorm.Model
	DirectURL   string `gorm:"uniqueIndex"`
	Size        int
	Width       int
	Height      int
	R18         bool
	Description string
	Extension   string
}

type Tag struct {
	gorm.Model
	Name     string
	Artworks []*Artwork `gorm:"many2many:artwork_tags;"`
}

type Artwork struct {
	gorm.Model
	Title     string
	Author    string
	Source    SourceName
	SourceURL string `gorm:"uniqueIndex"`
	Tags      []*Tag `gorm:"many2many:artwork_tags;"`
	Pictures  []*Picture
}

// 入库前的结构
type ArtworkRaw struct {
	Title     string
	Author    string
	Source    SourceName
	SourceURL string
	Tags      []string
	Pictures  []PictureRaw
}

type PictureRaw struct {
	DirectURL   string
	Size        int
	Width       int
	Height      int
	R18         bool
	Description string
	Extension   string
}
