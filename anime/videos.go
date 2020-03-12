package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetVideos(m common.MALItem) (Videos, error) {
	request, err := utils.CachedReqGet(utils.Constants.AppendAPIf("/anime/%v/videos", m.GetID()))
	if err != nil {
		return Videos{}, err
	}

	var a Videos
	err = request.ToJSON(&a)

	return a, err
}

type Videos struct {
	RequestHash        string  `json:"request_hash"`
	RequestCached      bool    `json:"request_cached"`
	RequestCacheExpiry int64   `json:"request_cache_expiry"`
	Promo              []Promo `json:"promo"`
}

type Promo struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
}
