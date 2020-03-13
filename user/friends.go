package user

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetFriends(m common.MALItem, page int) (Friends, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/user/%v/friends/%v", m.GetID(), page))
	if err != nil {
		return Friends{}, err
	}

	var a Friends
	err = request.ToJSON(&a)

	return a, err
}

type Friends struct {
	RequestHash        string   `json:"request_hash"`
	RequestCached      bool     `json:"request_cached"`
	RequestCacheExpiry int64    `json:"request_cache_expiry"`
	Friends            []Friend `json:"friends"`
}

type Friend struct {
	URL          string  `json:"url"`
	Username     string  `json:"username"`
	ImageURL     string  `json:"image_url"`
	LastOnline   string  `json:"last_online"`
	FriendsSince *string `json:"friends_since"`
}

func (r Friend) GetID() interface{} {
	return r.Username
}

func (r Friend) GetType() string {
	return mal_types.User
}
