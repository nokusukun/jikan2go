package magazine

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetMagazine(m common.MALItem, page int) (Magazine, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/magazine/%v/%v", m.GetID(), page))
	if err != nil {
		return Magazine{}, err
	}

	var a Magazine
	err = request.ToJSON(&a)

	return a, err
}

type Magazine struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	Meta               Meta           `json:"meta"`
	Manga              []MangaElement `json:"manga"`
}

type MangaElement struct {
	MalID           int64    `json:"mal_id"`
	URL             string   `json:"url"`
	Title           string   `json:"title"`
	ImageURL        string   `json:"image_url"`
	Synopsis        string   `json:"synopsis"`
	Type            string   `json:"type"`
	PublishingStart *string  `json:"publishing_start"`
	Volumes         *int64   `json:"volumes"`
	Members         int64    `json:"members"`
	Genres          []Meta   `json:"genres"`
	Authors         []Meta   `json:"authors"`
	Score           *float64 `json:"score"`
	Serialization   []string `json:"serialization"`
}

func (r MangaElement) GetID() interface{} {
	return r.MalID
}

func (r MangaElement) GetType() string {
	return mal_types.Manga
}

type Meta struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r Meta) GetID() interface{} {
	return r.MalID
}

func (r Meta) GetType() string {
	return "type"
}
