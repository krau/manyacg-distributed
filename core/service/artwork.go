package service

import (
	"github.com/krau/Picture-collector/core/dao"
	"github.com/krau/Picture-collector/core/models"
)

func AddArtworks(artworks []*models.ArtworkRaw) {
	artworkModels := make([]*models.Artwork, len(artworks))
	for i, artwork := range artworks {
		artworkModels[i] = artwork.ToArtwork()
	}
	dao.AddArtworks(artworkModels)
	for _, artwork := range artworks {
		artwork.ID = dao.GetArtworkBySourceURL(artwork.SourceURL).ID
	}
}
