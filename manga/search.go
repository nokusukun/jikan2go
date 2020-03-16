package manga

import (
	"encoding/json"
	"log"

	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func Search(q Query) (SearchResult, error) {
	if q.Page == 0 {
		q.Page = 1
	}

	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/search/manga"), q.ToParam())
	if err != nil {
		return SearchResult{}, err
	}

	var a SearchResult
	err = request.ToJSON(&a)

	return a, err
}

type Query struct {
	Q            string     `json:"q,omitempty"`
	Page         int        `json:"page,omitempty"`
	Type         Type       `json:"type,omitempty"`
	Rated        Rated      `json:"rated,omitempty"`
	Genre        MangaGenre `json:"genre,omitempty"`
	Score        float64    `json:"score,omitempty"`
	GenreExclude bool       `json:"genre_exclude,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Order        Order      `json:"order,omitempty"`
	Sort         Sort       `json:"sort,omitempty"`
	Producer     int64      `json:"producer,omitempty"`
	Magazine     int64      `json:"magazine,omitempty"`
	Letter       string     `json:"letter,omitempty"`
}

func (q *Query) ToParam() req.Param {
	qbytes, _ := json.Marshal(q)
	var reqParam req.Param
	err := json.Unmarshal(qbytes, &reqParam)
	if err != nil {
		log.Printf("failed to paramaterize %+v", q)
	}
	return reqParam
}

type SearchResult struct {
	RequestHash        string   `json:"request_hash"`
	RequestCached      bool     `json:"request_cached"`
	RequestCacheExpiry int64    `json:"request_cache_expiry"`
	Results            []Result `json:"results"`
	LastPage           int64    `json:"last_page"`
}

type Result struct {
	MalID      int64   `json:"mal_id"`
	URL        string  `json:"url"`
	ImageURL   string  `json:"image_url"`
	Title      string  `json:"title"`
	Publishing bool    `json:"publishing"`
	Synopsis   string  `json:"synopsis"`
	Type       Type    `json:"type"`
	Chapters   int64   `json:"chapters"`
	Volumes    int64   `json:"volumes"`
	Score      float64 `json:"score"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
	Members    int64   `json:"members"`
}

type Type string

const (
	TypeManga   Type = "Manga"
	TypeManhua  Type = "Manhua"
	TypeManhwa  Type = "Manhwa"
	TypeNovel   Type = "Novel"
	TypeOneShot Type = "One-shot"
)

type Rated string

const (
	G      Rated = "G"
	PG     Rated = "PG"
	PG13   Rated = "PG-13"
	R      Rated = "R"
	RatedR Rated = "R+"
	Rx     Rated = "Rx"
)

type Sort string

const (
	SortAsc  = "ascending"
	SortDesc = "descending"
)

type Order string

const (
	OrderTitle     = "title"
	OrderStartDate = "start_date"
	OrderEndDate   = "end_date"
	OrderScore     = "score"
	OrderType      = "type"
	OrderMemebers  = "members"
	OrderId        = "id"
	OrderEpisodes  = "episodes"
	OrderRating    = "rating"
)

type MangaGenre = int

//noinspection ALL,GoUnusedConst
const (
	GenreReserved MangaGenre = iota
	GenreAction
	GenreAdventure
	GenreCars
	GenreComedy
	GenreDementia
	GenreDemons
	GenreMystery
	GenreDrama
	GenreEcchi
	GenreFantasy
	GenreGame
	GenreHentai
	GenreHistorical
	GenreHorror
	GenreKids
	GenreMagic
	GenreMartialArts
	GenreMecha
	GenreMusic
	GenreParody
	GenreSamurai
	GenreRomance
	GenreSchool
	GenreSciFi
	GenreShoujo
	GenreShoujoAi
	GenreShounen
	GenreShounenAi
	GenreSpace
	GenreSports
	GenreSuperPower
	GenreVampire
	GenreYaoi
	GenreYuri
	GenreHarem
	GenreSliceOfLife
	GenreSupernatural
	GenreMilitary
	GenrePolice
	GenrePsychological
	GenreSeinen
	GenreJosei
	GenreDoujinshi
	GenreGenderBender
	GenreThriller
)
