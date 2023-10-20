package models

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	ArtworkID  string
	DirectURL  string  `gorm:"unique"`
	Hash       string  `gorm:"default:null"`
	BlurScore  float64 `gorm:"default:null"`
	Width      int     `gorm:"default:null"`
	Height     int     `gorm:"default:null"`
	Binary     []byte  `gorm:"default:null"`
	Downloaded bool    `gorm:"default:false"`
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
