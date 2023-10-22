package pixiv

import (
	"strings"
	"sync"

	"encoding/xml"

	"github.com/krau/manyacg/collector/common"
	"github.com/krau/manyacg/collector/config"
	"github.com/krau/manyacg/collector/logger"
	sourceModels "github.com/krau/manyacg/collector/sources/models"
	coreModels "github.com/krau/manyacg/core/models"
)

type SourcePixiv struct{}

func (sp *SourcePixiv) GetNewArtworks(limit int) ([]*coreModels.ArtworkRaw, error) {
	url := config.Cfg.Sources.Pixiv.URL
	logger.L.Infof("Fetching %s", url)
	resp, err := common.Cilent.R().Get(url)

	if err != nil {
		return nil, err
	}

	var pixivRss *sourceModels.PixivRss
	err = xml.NewDecoder(strings.NewReader(resp.String())).Decode(&pixivRss)

	if err != nil {
		return nil, err
	}

	logger.L.Debugf("Got %d items", len(pixivRss.Channel.Items))
	artworkChan := make(chan *coreModels.ArtworkRaw)
	var wg sync.WaitGroup

	for i, item := range pixivRss.Channel.Items {
		if i >= limit {
			break
		}

		wg.Add(1)
		go func(item sourceModels.Item) {
			defer wg.Done()

			imgs := strings.Split(item.Description, "<img src=\"")
			srcs := make([]string, 0)
			for _, img := range imgs {
				if strings.HasPrefix(img, "http") {
					src := strings.Split(img, "\"")[0]
					srcs = append(srcs, src)
				}
			}
			pictures := make([]*coreModels.PictureRaw, 0)
			for _, src := range srcs {
				picture := coreModels.PictureRaw{
					DirectURL: src,
				}
				pictures = append(pictures, &picture)
			}

			artworkInfo, err := getArtworkInfo(item.Link)
			if err != nil {
				logger.L.Errorf("Error fetching artwork info: %v", err)
				return
			}

			tags := make([]string, 0)
			for _, tag := range artworkInfo.Body.Tags.Tags {
				var tagName string
				if tag.Translation != nil {
					tagName = tag.Translation.En
				} else {
					tagName = tag.Tag
				}
				tags = append(tags, tagName)
			}
			isR18 := false
			for _, tag := range tags {
				switch tag {
				case "R-18":
					isR18 = true
					break
				case "R-18G":
					isR18 = true
					break
				case "R18":
					isR18 = true
					break
				case "R18G":
					isR18 = true
					break
				}
			}
			artwork := coreModels.ArtworkRaw{
				Title:       item.Title,
				Author:      item.Author,
				Description: artworkInfo.Body.ExtraData.Meta.Description,
				Source:      coreModels.SourcePixiv,
				SourceURL:   item.Link,
				Tags:        tags,
				R18:         isR18,
				Pictures:    pictures,
			}
			artworkChan <- &artwork
		}(item)
	}

	go func() {
		wg.Wait()
		close(artworkChan)
	}()

	artworks := make([]*coreModels.ArtworkRaw, 0)
	for artwork := range artworkChan {
		artworks = append(artworks, artwork)
	}

	return artworks, nil
}

func (sp *SourcePixiv) SourceName() coreModels.SourceName {
	return coreModels.SourcePixiv
}

func (sp *SourcePixiv) Config() *config.SourceConfig {
	return &config.Cfg.Sources.Pixiv
}
