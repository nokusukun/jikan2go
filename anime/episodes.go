package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetEpisodes(anime common.MALItem) (Episodes, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/anime/%v/episodes", anime.GetID()))
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
