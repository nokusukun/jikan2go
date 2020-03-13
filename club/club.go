package club

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

// https://api.jikan.moe/v3/club/1/

func GetClub(club common.MALItem) (Club, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/club/%v", club.GetID()))
	if err != nil {
		return Club{}, err
	}

	var a Club
	err = request.ToJSON(&a)

	return a, err
}

func GetMembers(club common.MALItem, page int) (Members, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/club/%v/members/%v", club.GetID(), page))
	if err != nil {
		return Members{}, err
	}

	var a Members
	err = request.ToJSON(&a)

	return a, err
}

type Club struct {
	RequestHash        string             `json:"request_hash"`
	RequestCached      bool               `json:"request_cached"`
	RequestCacheExpiry int64              `json:"request_cache_expiry"`
	MalID              int64              `json:"mal_id"`
	URL                string             `json:"url"`
	ImageURL           string             `json:"image_url"`
	Title              string             `json:"title"`
	MembersCount       int64              `json:"members_count"`
	PicturesCount      int64              `json:"pictures_count"`
	Category           string             `json:"category"`
	Created            string             `json:"created"`
	Type               string             `json:"type"`
	Staff              []common.User      `json:"staff"`
	AnimeRelations     []common.Anime     `json:"anime_relations"`
	MangaRelations     []common.Manga     `json:"manga_relations"`
	CharacterRelations []common.Character `json:"character_relations"`
}

func (r Club) GetID() interface{} {
	return r.MalID
}

func (r Club) GetType() string {
	return mal_types.Club
}

type Members struct {
	RequestHash        string          `json:"request_hash"`
	RequestCached      bool            `json:"request_cached"`
	RequestCacheExpiry int64           `json:"request_cache_expiry"`
	Members            []common.Member `json:"members"`
}
