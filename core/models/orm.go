package models

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	ArtworkID  uint    `gorm:"index"`
	DirectURL  string  `gorm:"unique"`
	Hash       string  `gorm:"default:null"`
	BlurScore  float64 `gorm:"default:null"`
	Width      uint    `gorm:"default:null"`
	Height     uint    `gorm:"default:null"`
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
	SourceURL   string    `gorm:"unique"`
	Tags        []*Tag    `gorm:"many2many:artwork_tags;"`
	R18         bool      `gorm:"default:false"`
	Pictures    []Picture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Tag) String() string {
	return t.Name
}
