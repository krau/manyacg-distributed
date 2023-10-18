package models

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	ArtworkID string
	DirectURL string `gorm:"unique"`
}

type Tag struct {
	gorm.Model
	Name     string     `gorm:"unique"`
	Artworks []*Artwork `gorm:"many2many:artwork_tags;"`
}

type Artwork struct {
	gorm.Model
	Title       string
	Author      string
	Description string
	Source      SourceName
	SourceURL   string `gorm:"unique"`
	Tags        []*Tag `gorm:"many2many:artwork_tags;"`
	R18         bool   `gorm:"default:false"`
	Pictures    []Picture
}
