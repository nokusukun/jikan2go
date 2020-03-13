package common

import (
	"github.com/nokusukun/jikan2go/utils"
)

func GetInfo(m MALItem) (Info, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/%v/%v/moreinfo", m.GetType(), m.GetID()))
	if err != nil {
		return Info{}, err
	}

	var a Info
	err = request.ToJSON(&a)

	return a, err
}

type Info struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int64  `json:"request_cache_expiry"`
	Info               string `json:"moreinfo"`
}
