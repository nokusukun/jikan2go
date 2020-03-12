package common

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func GetPictures(m MALItem) (Pictures, error) {
	request, err := req.Get(utils.Constants.AppendAPIf("/%v/%v/pictures", m.GetType(), m.GetID()))
	if err != nil {
		return Pictures{}, err
	}

	var a Pictures
	err = request.ToJSON(&a)

	return a, err
}

type Pictures struct {
	RequestHash        string    `json:"request_hash"`
	RequestCached      bool      `json:"request_cached"`
	RequestCacheExpiry int64     `json:"request_cache_expiry"`
	Pictures           []Picture `json:"pictures"`
}

type Picture struct {
	Large string `json:"large"`
	Small string `json:"small"`
}
