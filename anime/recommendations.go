package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetRecommendations(anime common.MALItem) (Recommendations, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/anime/%v/recommendations", anime.GetID()))
	if err != nil {
		return Recommendations{}, err
	}

	var a Recommendations
	err = request.ToJSON(&a)

	return a, err
}

type Recommendations struct {
	RequestHash        string           `json:"request_hash"`
	RequestCached      bool             `json:"request_cached"`
	RequestCacheExpiry int64            `json:"request_cache_expiry"`
	Recommendations    []Recommendation `json:"recommendations"`
}

type Recommendation struct {
	MalID               int64  `json:"mal_id"`
	URL                 string `json:"url"`
	ImageURL            string `json:"image_url"`
	RecommendationURL   string `json:"recommendation_url"`
	Title               string `json:"title"`
	RecommendationCount int64  `json:"recommendation_count"`
}

func (r Recommendation) GetID() interface{} {
	return r.MalID
}

func (r Recommendation) GetType() string {
	return mal_types.Anime
}
