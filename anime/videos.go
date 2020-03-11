package anime

import (
    "github.com/imroc/req"

    "github.com/nokusukun/Jikan2Go/common"
    "github.com/nokusukun/Jikan2Go/utils"
)

func GetVideos(m common.MALItem) (Videos, error) {
    request, err := req.Get(utils.Contstants.AppendAPIf("/anime/%v/videos", m.GetID()))
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
}

type Promo struct {
    Title    string `json:"title"`
    ImageURL string `json:"image_url"`
    VideoURL string `json:"video_url"`
}
