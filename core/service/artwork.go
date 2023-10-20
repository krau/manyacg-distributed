package service

import (
	"github.com/krau/Picture-collector/core/dao"
	"github.com/krau/Picture-collector/core/models"
)

func AddArtworks(artworks []*models.ArtworkRaw) []*models.ArtworkRaw {
	existingArtworks := make(map[string]*models.Artwork)
	for _, artwork := range artworks {
		existingArtwork := dao.GetArtworkBySourceURL(artwork.SourceURL)
		if existingArtwork != nil {
			existingArtworks[artwork.SourceURL] = existingArtwork
		}
	}

	artworkModels := make([]*models.Artwork, len(artworks))
	newArtworks := make([]*models.ArtworkRaw, 0)
	for i, artwork := range artworks {
		artworkModels[i] = artwork.ToArtwork()
		if existingArtworks[artwork.SourceURL] == nil {
			newArtworks = append(newArtworks, artwork)
		}
	}
	dao.AddArtworks(artworkModels)

	for _, artwork := range artworks {
		if existingArtwork := existingArtworks[artwork.SourceURL]; existingArtwork != nil {
			artwork.ID = existingArtwork.ID
		}
	}

	return newArtworks
}
