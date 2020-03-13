package person

import (
	"encoding/json"
	"log"

	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func Search(q Query) (SearchResult, error) {
	if q.Page == 0 {
		q.Page = 1
	}

	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/search/people"), q.ToParam())
	if err != nil {
		return SearchResult{}, err
	}

	var a SearchResult
	err = request.ToJSON(&a)

	return a, err
}

type Query struct {
	Q            string  `json:"q,omitempty"`
	Page         int     `json:"page,omitempty"`
	Score        float64 `json:"score,omitempty"`
	GenreExclude bool    `json:"genre_exclude,omitempty"`
	Limit        int     `json:"limit,omitempty"`
	Producer     int64   `json:"producer,omitempty"`
	Magazine     int64   `json:"magazine,omitempty"`
	Letter       string  `json:"letter,omitempty"`
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
	MalID            int64    `json:"mal_id"`
	URL              string   `json:"url"`
	ImageURL         string   `json:"image_url"`
	Name             string   `json:"name"`
	AlternativeNames []string `json:"alternative_names"`
}

func (r Result) GetID() interface{} {
	return r.MalID
}

func (r Result) GetType() string {
	return mal_types.Person
}
