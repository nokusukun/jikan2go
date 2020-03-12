package character

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func GetTop(page int) (Top, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/top/people/%v", page))
	if err != nil {
		return Top{}, err
	}

	var a Top
	err = request.ToJSON(&a)

	return a, err
}

type Top struct {
	RequestHash        string       `json:"request_hash"`
	RequestCached      bool         `json:"request_cached"`
	RequestCacheExpiry int64        `json:"request_cache_expiry"`
	Top                []TopElement `json:"top"`
}

type TopElement struct {
	MalID        int64     `json:"mal_id"`
	Rank         int64     `json:"rank"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	NameKanji    string    `json:"name_kanji"`
	Animeography []Ography `json:"animeography"`
	Mangaography []Ography `json:"mangaography"`
	Favorites    int64     `json:"favorites"`
	ImageURL     string    `json:"image_url"`
}

func (r TopElement) GetID() int64 {
	return r.MalID
}

func (r TopElement) GetType() string {
	return "character"
}

type Ography struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r Ography) GetID() int64 {
	return r.MalID
}

func (r Ography) GetType() string {
	return r.Type
}
