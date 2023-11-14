package entity

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string     `gorm:"unique"`
	Artworks []*Artwork `gorm:"many2many:artwork_tags;"`
}

func (t *Tag) String() string {
	return t.Name
}
