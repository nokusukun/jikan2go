// anime contains several key methods and structs for interacting with Jikan's resources and endpoints.
package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

// GetAnime returns a canonical Anime Object
// Return anime with given an interface that has a mal_id.
// This method parses item data by ID from https://myanimelist.net/anime/{id}
func GetAnime(anime common.MALItem) (Anime, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/anime/%v", anime.GetID()))
	if err != nil {
		return Anime{}, err
	}

	var a Anime
	err = request.ToJSON(&a)

	return a, err
}

// A single anime object with all its canonical details
type Anime struct {
	RequestHash        string                `json:"request_hash"`
	RequestCached      bool                  `json:"request_cached"`
	RequestCacheExpiry int64                 `json:"request_cache_expiry"`
	MalID              int64                 `json:"mal_id"`
	URL                string                `json:"url"`
	ImageURL           string                `json:"image_url"`
	TrailerURL         string                `json:"trailer_url"`
	Title              string                `json:"title"`
	TitleEnglish       string                `json:"title_english"`
	TitleJapanese      string                `json:"title_japanese"`
	TitleSynonyms      []string              `json:"title_synonyms"`
	Type               string                `json:"type"`
	Source             string                `json:"source"`
	Episodes           int64                 `json:"episodes"`
	Status             string                `json:"status"`
	Airing             bool                  `json:"airing"`
	Aired              Aired                 `json:"aired"`
	Duration           string                `json:"duration"`
	Rating             string                `json:"rating"`
	Score              float64               `json:"score"`
	ScoredBy           int64                 `json:"scored_by"`
	Rank               int64                 `json:"rank"`
	Popularity         int64                 `json:"popularity"`
	Members            int64                 `json:"members"`
	Favorites          int64                 `json:"favorites"`
	Synopsis           string                `json:"synopsis"`
	Background         string                `json:"background"`
	Premiered          string                `json:"premiered"`
	Broadcast          string                `json:"broadcast"`
	Related            Related               `json:"related"`
	Producers          []common.TypedMALItem `json:"producers"`
	Licensors          []common.TypedMALItem `json:"licensors"`
	Studios            []common.TypedMALItem `json:"studios"`
	Genres             []common.TypedMALItem `json:"genres"`
	OpeningThemes      []string              `json:"opening_themes"`
	EndingThemes       []string              `json:"ending_themes"`
}

type Aired struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Prop   Prop   `json:"prop"`
	String string `json:"string"`
}

type Prop struct {
	From Date `json:"from"`
	To   Date `json:"to"`
}

type Date struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

type Related struct {
	Adaptation []common.TypedMALItem `json:"Adaptation"`
	SideStory  []common.TypedMALItem `json:"Side story"`
	Summary    []common.TypedMALItem `json:"Summary"`
}

//noinspection ALL
const (
	TManga Type = "manga"
	TAnime Type = "anime"
)

func (r Anime) GetID() interface{} {
	return r.MalID
}

func (r Anime) GetType() string {
	return mal_types.Anime
}
