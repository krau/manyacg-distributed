package models

import "gorm.io/gorm"

type Picture struct {
	gorm.Model
	DirectURL string `gorm:"uniqueIndex"`
}

type Tag struct {
	gorm.Model
	Name     string
	Artworks []*Artwork `gorm:"many2many:artwork_tags;"`
}

type Artwork struct {
	gorm.Model
	Title       string
	Author      string
	Description string
	Source      SourceName
	SourceURL   string `gorm:"uniqueIndex"`
	Tags        []*Tag `gorm:"many2many:artwork_tags;"`
	R18         bool   `gorm:"default:false"`
	Pictures    []*Picture
}

// 入库前的结构
type ArtworkRaw struct {
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
	DirectURL string
}
