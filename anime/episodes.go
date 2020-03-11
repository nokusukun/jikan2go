package anime

import (
    "github.com/imroc/req"

    "github.com/nokusukun/Jikan2Go/common"
    "github.com/nokusukun/Jikan2Go/utils"
)

func GetEpisodes(m common.MALItem) (Episodes, error) {
    request, err := req.Get(utils.Contstants.AppendAPIf("/anime/%v/episodes", m.GetID()))
    if err != nil {
        return Episodes{}, err
    }

    var a Episodes
    err = request.ToJSON(&a)

    return a, err
}


type Episodes struct {
    RequestHash        string    `json:"request_hash"`
    RequestCached      bool      `json:"request_cached"`
    RequestCacheExpiry int64     `json:"request_cache_expiry"`
    EpisodesLastPage   int64     `json:"episodes_last_page"`
    Episodes           []Episode `json:"episodes"`
}

type Episode struct {
    EpisodeID     int64       `json:"episode_id"`
    Title         string      `json:"title"`
    TitleJapanese string      `json:"title_japanese"`
    TitleRomanji  string      `json:"title_romanji"`
    Aired         string      `json:"aired"`
    Filler        bool        `json:"filler"`
    Recap         bool        `json:"recap"`
    VideoURL      interface{} `json:"video_url"`
    ForumURL      string      `json:"forum_url"`
}
