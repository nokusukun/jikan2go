package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetVideos(anime common.MALItem) (Videos, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/anime/%v/videos", anime.GetID()))
	if err != nil {
		return Videos{}, err
	}

	var a Videos
	err = request.ToJSON(&a)

	return a, err
}

type Videos struct {
	RequestHash        string        `json:"request_hash"`
	RequestCached      bool          `json:"request_cached"`
	RequestCacheExpiry int64         `json:"request_cache_expiry"`
	Promo              []Promo       `json:"promo"`
	Episodes           []EpisodeInfo `json:"episodes"`
}

type EpisodeInfo struct {
	Title    string `json:"title"`
	Episode  string `json:"episode"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

type Promo struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
}
