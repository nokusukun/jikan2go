// studio implements Jikan's api/studio resources and endpoints
package studio

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/genre"
	"github.com/nokusukun/jikan2go/utils"
)

// Get studio retrieves the canonical data of a Studio, ensure that the MALItem you're using as an argument
// is a studio, or else incorrect data may be returned.
func GetStudio(m common.MALItem, page int) (Studio, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/producer/%v/%v", m.GetID(), page))
	if err != nil {
		return Studio{}, err
	}

	var a Studio
	err = request.ToJSON(&a)

	return a, err
}

type Studio struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	Meta               StudioMeta     `json:"meta"`
	Anime              []AnimeElement `json:"anime"`
}

type AnimeElement struct {
	MalID       int64         `json:"mal_id"`
	URL         string        `json:"url"`
	Title       string        `json:"title"`
	ImageURL    string        `json:"image_url"`
	Synopsis    string        `json:"synopsis"`
	Type        AnimeType     `json:"type"`
	AiringStart *string       `json:"airing_start"`
	Episodes    *int64        `json:"episodes"`
	Members     int64         `json:"members"`
	Genres      []genre.Genre `json:"genres"`
	Source      Source        `json:"source"`
	Producers   []StudioMeta  `json:"producers"`
	Score       float64       `json:"score"`
	Licensors   []string      `json:"licensors"`
	R18         bool          `json:"r18"`
	Kids        bool          `json:"kids"`
}

func (r AnimeElement) GetID() int64 {
	return r.MalID
}

func (r AnimeElement) GetType() string {
	return "anime"
}

type StudioMeta struct {
	MalID int64    `json:"mal_id"`
	Type  MetaType `json:"type"`
	Name  string   `json:"name"`
	URL   string   `json:"url"`
}

func (r StudioMeta) GetID() int64 {
	return r.MalID
}

func (r StudioMeta) GetType() string {
	return "studio"
}

type MetaType string

const (
	Anime MetaType = "anime"
)

type Source string

const (
	Empty      Source = "-"
	Game       Source = "Game"
	LightNovel Source = "Light novel"
	Manga      Source = "Manga"
	Music      Source = "Music"
	Novel      Source = "Novel"
	Original   Source = "Original"
)

type AnimeType string

const (
	Movie   AnimeType = "Movie"
	Ona     AnimeType = "ONA"
	Ova     AnimeType = "OVA"
	Special AnimeType = "Special"
	Tv      AnimeType = "TV"
)
