package genre

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/manga"
	"github.com/nokusukun/jikan2go/person"
	"github.com/nokusukun/jikan2go/utils"
)

func GetManga(genre manga.MangaGenre, page int) (MangaGenre, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/genre/manga/%v/%v", genre, page))
	if err != nil {
		return MangaGenre{}, err
	}

	var a MangaGenre
	err = request.ToJSON(&a)

	return a, err
}

type MangaGenre struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	MalURL             Genre          `json:"mal_url"`
	ItemCount          int64          `json:"item_count"`
	Manga              []MangaElement `json:"manga"`
}

type MangaElement struct {
	MalID           int64           `json:"mal_id"`
	URL             string          `json:"url"`
	Title           string          `json:"title"`
	ImageURL        string          `json:"image_url"`
	Synopsis        string          `json:"synopsis"`
	Type            manga.Type      `json:"type"`
	PublishingStart *string         `json:"publishing_start"`
	Volumes         *int64          `json:"volumes"`
	Members         int64           `json:"members"`
	Genres          []Genre         `json:"genres"`
	Authors         []person.Person `json:"authors"`
	Score           float64         `json:"score"`
	Serialization   []string        `json:"serialization"`
}

func (r MangaElement) GetID() int64 {
	return r.MalID
}

func (r MangaElement) GetType() string {
	return "manga"
}
