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
	FilePath   string  `gorm:"default:null"`
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

func (p *Picture) ToResp() *RespPicture {
	return &RespPicture{
		Status:    0,
		Message:   "success",
		CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
		DirectURL: p.DirectURL,
		Width:     p.Width,
		Height:    p.Height,
		BlurScore: p.BlurScore,
		Hash:      p.Hash,
	}
}

func (a *Artwork) ToResp() *RespArtwork {
	tags := make([]string, len(a.Tags))
	for i, tag := range a.Tags {
		tags[i] = tag.String()
	}

	pictures := make([]RespPicture, len(a.Pictures))
	for i, picture := range a.Pictures {
		pictures[i] = *picture.ToResp()
	}

	return &RespArtwork{
		Status:      0,
		Message:     "success",
		CreatedAt:   a.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   a.UpdatedAt.Format("2006-01-02 15:04:05"),
		Title:       a.Title,
		Author:      a.Author,
		Description: a.Description,
		Source:      a.Source.String(),
		SourceURL:   a.SourceURL,
		Tags:        tags,
		R18:         a.R18,
		Pictures:    pictures,
	}
}
