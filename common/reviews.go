package common

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func GetReviews(m MALItem) (Reviews, error) {
	request, err := req.Get(utils.Constants.AppendAPIf("/%v/%v/reviews", m.GetType(), m.GetID()))
	if err != nil {
		return Reviews{}, err
	}

	var a Reviews
	err = request.ToJSON(&a)

	return a, err
}

type Reviews struct {
	RequestHash        string   `json:"request_hash"`
	RequestCached      bool     `json:"request_cached"`
	RequestCacheExpiry int64    `json:"request_cache_expiry"`
	Reviews            []Review `json:"reviews"`
}

type Review struct {
	MalID        int64       `json:"mal_id"`
	URL          string      `json:"url"`
	Type         interface{} `json:"type"`
	HelpfulCount int64       `json:"helpful_count"`
	Date         string      `json:"date"`
	Reviewer     Reviewer    `json:"reviewer"`
	Content      string      `json:"content"`
}

type Reviewer struct {
	URL          string `json:"url"`
	ImageURL     string `json:"image_url"`
	Username     string `json:"username"`
	EpisodesSeen int64  `json:"episodes_seen"`
	Scores       Scores `json:"scores"`
}

type Scores struct {
	Overall   int64 `json:"overall"`
	Story     int64 `json:"story"`
	Animation int64 `json:"animation"`
	Sound     int64 `json:"sound"`
	Character int64 `json:"character"`
	Enjoyment int64 `json:"enjoyment"`
}
