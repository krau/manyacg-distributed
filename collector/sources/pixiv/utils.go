package pixiv

import (
	"encoding/json"
	"strings"

	"github.com/krau/Picture-collector/collector/common"
	"github.com/krau/Picture-collector/collector/logger"
	sourceModels "github.com/krau/Picture-collector/collector/sources/models"
)

func getArtworkInfo(sourceURL string) (*sourceModels.PixivAjaxResp, error) {
	pid := strings.Split(sourceURL, "/")[len(strings.Split(sourceURL, "/"))-1]
	ajaxURL := "https://www.pixiv.net/ajax/illust/" + pid
	logger.L.Debugf("Fetching artwork info: %s", ajaxURL)
	resp, err := common.Cilent.R().Get(ajaxURL)
	if err != nil {
		return nil, err
	}
	var pixivAjaxResp sourceModels.PixivAjaxResp
	err = json.Unmarshal([]byte(resp.String()), &pixivAjaxResp)
	if err != nil {
		return nil, err
	}
	return &pixivAjaxResp, nil
}