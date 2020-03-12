package anime

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func GetTop(page int, subtype SubType) (Top, error) {
	request, err := req.Get(utils.Contstants.AppendAPIf("/top/anime/%v/%v", page, subtype))
	if err != nil {
		return Top{}, err
	}

	var a Top
	err = request.ToJSON(&a)

	return a, err
}

type SubType string

const (
	SubTypeAiring     SubType = "airing"
	SubTypeUpcoming   SubType = "upcoming"
	SubTypeTv         SubType = "tv"
	SubTypeMovie      SubType = "movie"
	SubTypeOVA        SubType = "ova"
	SubTypeSpecial    SubType = "special"
	SubTypePopularity SubType = "bypopularity"
	SubTypeFavourite  SubType = "favorite"
)

type Top struct {
	RequestHash        string       `json:"request_hash"`
	RequestCached      bool         `json:"request_cached"`
	RequestCacheExpiry int64        `json:"request_cache_expiry"`
	Top                []TopElement `json:"top"`
}

type TopElement struct {
	MalID     int64   `json:"mal_id"`
	Rank      int64   `json:"rank"`
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	ImageURL  string  `json:"image_url"`
	Type      Type    `json:"type"`
	Episodes  int64   `json:"episodes"`
	StartDate string  `json:"start_date"`
	EndDate   *string `json:"end_date"`
	Members   int64   `json:"members"`
	Score     float64 `json:"score"`
}

func (r TopElement) GetID() int64 {
	return r.MalID
}

func (r TopElement) GetType() string {
	return "anime"
}
