package manga

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetCharacters(m common.MALItem) (Characters, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/manga/%v/characters", m.GetID()))
	if err != nil {
		return Characters{}, err
	}

	var a Characters
	err = request.ToJSON(&a)

	return a, err
}

type Characters struct {
	RequestHash        string      `json:"request_hash"`
	RequestCached      bool        `json:"request_cached"`
	RequestCacheExpiry int64       `json:"request_cache_expiry"`
	Characters         []Character `json:"characters"`
}

type Character struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Role     Role   `json:"role"`
}

type Role string

const (
	Main       Role = "Main"
	Supporting Role = "Supporting"
)
