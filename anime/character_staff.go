package anime

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetCharacterStaff(anime common.MALItem) (CharacterStaff, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/anime/%v/characters_staff", anime.GetID()))
	if err != nil {
		return CharacterStaff{}, err
	}

	var a CharacterStaff
	err = request.ToJSON(&a)

	return a, err
}

type CharacterStaff struct {
	RequestHash        string      `json:"request_hash"`
	RequestCached      bool        `json:"request_cached"`
	RequestCacheExpiry int64       `json:"request_cache_expiry"`
	Characters         []Character `json:"characters"`
	Staff              []Staff     `json:"staff"`
}

type Character struct {
	MalID       int64   `json:"mal_id"`
	URL         string  `json:"url"`
	ImageURL    string  `json:"image_url"`
	Name        string  `json:"name"`
	Role        Role    `json:"role"`
	VoiceActors []Staff `json:"voice_actors"`
}

type Staff struct {
	MalID     int64     `json:"mal_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	ImageURL  string    `json:"image_url"`
	Language  *Language `json:"language,omitempty"`
	Positions []string  `json:"positions"`
}

type Role string

const (
	Main       Role = "Main"
	Supporting Role = "Supporting"
)

type Language string

const (
	English  Language = "English"
	German   Language = "German"
	Japanese Language = "Japanese"
)

func (r Character) GetID() interface{} {
	return r.MalID
}

func (r Character) GetType() string {
	return mal_types.Character
}

func (r Staff) GetID() interface{} {
	return r.MalID
}

func (r Staff) GetType() string {
	return mal_types.Person
}
