// manga implements Jikan's api/manga resources and endpoints
package manga

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetManga(m common.MALItem) (Manga, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/manga/%v", m.GetID()))
	if err != nil {
		return Manga{}, err
	}

	var a Manga
	err = request.ToJSON(&a)

	return a, err
}

type Manga struct {
	RequestHash        string    `json:"request_hash"`
	RequestCached      bool      `json:"request_cached"`
	RequestCacheExpiry int64     `json:"request_cache_expiry"`
	MalID              int64     `json:"mal_id"`
	URL                string    `json:"url"`
	Title              string    `json:"title"`
	TitleEnglish       string    `json:"title_english"`
	TitleSynonyms      []string  `json:"title_synonyms"`
	TitleJapanese      string    `json:"title_japanese"`
	Status             string    `json:"status"`
	ImageURL           string    `json:"image_url"`
	Type               string    `json:"type"`
	Volumes            int64     `json:"volumes"`
	Chapters           int64     `json:"chapters"`
	Publishing         bool      `json:"publishing"`
	Published          Published `json:"published"`
	Rank               int64     `json:"rank"`
	Score              float64   `json:"score"`
	ScoredBy           int64     `json:"scored_by"`
	Popularity         int64     `json:"popularity"`
	Members            int64     `json:"members"`
	Favorites          int64     `json:"favorites"`
	Synopsis           string    `json:"synopsis"`
	Background         string    `json:"background"`
	Related            Related   `json:"related"`
	Genres             []MALItem `json:"genres"`
	Authors            []MALItem `json:"authors"`
	Serializations     []MALItem `json:"serializations"`
}

func (r Manga) GetID() int64 {
	return r.MalID
}

func (r Manga) GetType() string {
	return "manga"
}

type MALItem struct {
	MalID int64  `json:"mal_id"`
	Type  Type   `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r MALItem) GetID() int64 {
	return r.MalID
}

func (r MALItem) GetType() string {
	return "person"
}

type Published struct {
	From   string      `json:"from"`
	To     interface{} `json:"to"`
	Prop   Prop        `json:"prop"`
	String string      `json:"string"`
}

type Prop struct {
	From From `json:"from"`
	To   From `json:"to"`
}

type From struct {
	Day   *int64 `json:"day"`
	Month *int64 `json:"month"`
	Year  *int64 `json:"year"`
}

type Related struct {
	Other              []MALItem `json:"Other"`
	AlternativeVersion []MALItem `json:"Alternative version"`
	SideStory          []MALItem `json:"Side story"`
	SpinOff            []MALItem `json:"Spin-off"`
	Adaptation         []MALItem `json:"Adaptation"`
}

const (
	TypeAnime  Type = "anime"
	TypePeople Type = "people"
)
