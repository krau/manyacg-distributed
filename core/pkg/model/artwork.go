package model

import "github.com/krau/manyacg/core/pkg/common/enum/source"

type ArtworkRaw struct {
	ID          uint
	Title       string
	Author      string
	Description string
	Source      source.SourceName
	SourceURL   string
	Tags        []string
	R18         bool
	Pictures    []*PictureRaw
}

type ProcessedArtwork struct {
	ArtworkID   uint     `json:"artwork_id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Source      string   `json:"source"`
	SourceURL   string   `json:"source_url"`
	Tags        []string `json:"tags"`
	R18         bool     `json:"r18"`
	Pictures    []*struct {
		DirectURL string `json:"direct_url"`
	} `json:"pictures"`
}
