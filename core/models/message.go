package models

type MessageProcessedArtwork struct {
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
