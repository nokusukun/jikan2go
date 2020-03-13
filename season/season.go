// season implements Jikan's api/season resources and endpoints
package season

import (
	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

type S string

const (
	Winter S = "winter"
	Summer S = "summer"
	Fall   S = "fall"
	Spring S = "spring"
)

func GetSeasonLater() (Season, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/season/later"))
	if err != nil {
		return Season{}, err
	}

	var a Season
	err = request.ToJSON(&a)

	return a, err
}

func GetSeason(season S, year int) (Season, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/season/%v/%v", year, season))
	if err != nil {
		return Season{}, err
	}

	var a Season
	err = request.ToJSON(&a)

	return a, err
}

type Season struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	SeasonName         string         `json:"season_name"`
	SeasonYear         int64          `json:"season_year"`
	Anime              []AnimeElement `json:"anime"`
}

type AnimeElement struct {
	MalID       int64          `json:"mal_id"`
	URL         string         `json:"url"`
	Title       string         `json:"title"`
	ImageURL    string         `json:"image_url"`
	Synopsis    string         `json:"synopsis"`
	Type        anime.Type     `json:"type"`
	AiringStart string         `json:"airing_start"`
	Episodes    int64          `json:"episodes"`
	Members     int64          `json:"members"`
	Genres      []common.Genre `json:"genres"`
	Source      Source         `json:"source"`
	Producers   []common.Genre `json:"producers"`
	Score       float64        `json:"score"`
	Licensors   []string       `json:"licensors"`
	R18         bool           `json:"r18"`
	Kids        bool           `json:"kids"`
	Continuing  bool           `json:"continuing"`
}

func (r AnimeElement) GetID() interface{} {
	return r.MalID
}

func (r AnimeElement) GetType() string {
	return mal_types.Anime
}

type Source string

const (
	CardGame      Source = "Card game"
	Empty         Source = "-"
	Game          Source = "Game"
	LightNovel    Source = "Light novel"
	Manga         Source = "Manga"
	Novel         Source = "Novel"
	Original      Source = "Original"
	Other         Source = "Other"
	PictureBook   Source = "Picture book"
	The4KomaManga Source = "4-koma manga"
	VisualNovel   Source = "Visual novel"
	WebManga      Source = "Web manga"
)
