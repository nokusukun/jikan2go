package manga

import (
	"github.com/nokusukun/jikan2go/utils"
)

func GetTop(page int, subtype SubType) (Top, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/top/manga/%v/%v", page, subtype))
	if err != nil {
		return Top{}, err
	}

	var a Top
	err = request.ToJSON(&a)

	return a, err
}

type SubType string

// manga novels oneshots doujin manhwa manhua
const (
	SubTypeManga      SubType = "manga"
	SubTypeNovels     SubType = "novels"
	SubTypeOneshots   SubType = "oneshots"
	SubTypeDoujin     SubType = "doujin"
	SubTypeManhwa     SubType = "manhwa"
	SubTypeManhua     SubType = "manhua"
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
	Type      Type    `json:"type"`
	Volumes   *int64  `json:"volumes"`
	StartDate string  `json:"start_date"`
	EndDate   *string `json:"end_date"`
	Members   int64   `json:"members"`
	Score     float64 `json:"score"`
	ImageURL  string  `json:"image_url"`
}

func (r TopElement) GetID() int64 {
	return r.MalID
}

func (r TopElement) GetType() string {
	return "manga"
}

// manga novels oneshots doujin manhwa manhua
const (
	TopManga    Type = "Manga"
	TopManhwa   Type = "Manhwa"
	TopNovel    Type = "Novel"
	TopOneshots Type = "Oneshots"
	TopDoujin   Type = "Doujin"
	TopManhua   Type = "Manhua"
)
