package genre

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/utils"
)

func GetAnime(genre anime.AnimeGenre, page int) (AnimeGenre, error) {
	request, err := req.Get(utils.Constants.AppendAPIf("/genre/anime/%v/%v", genre, page))
	if err != nil {
		return AnimeGenre{}, err
	}

	var a AnimeGenre
	err = request.ToJSON(&a)

	return a, err
}

type AnimeGenre struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	MalURL             Genre          `json:"mal_url"`
	ItemCount          int64          `json:"item_count"`
	Anime              []AnimeElement `json:"anime"`
}

type AnimeElement struct {
	MalID       int64     `json:"mal_id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	ImageURL    string    `json:"image_url"`
	Synopsis    string    `json:"synopsis"`
	Type        AnimeType `json:"type"`
	AiringStart string    `json:"airing_start"`
	Episodes    *int64    `json:"episodes"`
	Members     int64     `json:"members"`
	Genres      []Genre   `json:"genres"`
	Source      Source    `json:"source"`
	Producers   []Studio  `json:"producers"`
	Score       float64   `json:"score"`
	Licensors   []string  `json:"licensors"`
	R18         bool      `json:"r18"`
	Kids        bool      `json:"kids"`
}

func (r AnimeElement) GetID() int64 {
	return r.MalID
}

func (r AnimeElement) GetType() string {
	return "anime"
}

type Genre struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r Genre) GetID() int64 {
	return r.MalID
}

func (r Genre) GetType() string {
	return "genre/" + r.Type
}

type Studio struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r Studio) GetID() int64 {
	return r.MalID
}

func (r Studio) GetType() string {
	return "studio"
}

type Source string

const (
	LightNovel  Source = "Light novel"
	Manga       Source = "Manga"
	Original    Source = "Original"
	VisualNovel Source = "Visual novel"
	WebManga    Source = "Web manga"
)

type AnimeType string

const (
	Movie AnimeType = "Movie"
	Ova   AnimeType = "OVA"
	Tv    AnimeType = "TV"
)
