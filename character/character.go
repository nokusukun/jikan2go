// character implements Jikan's api/character resources and endpoints
package character

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetCharacter(character common.MALItem) (Character, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/character/%v", character.GetID()))
	if err != nil {
		return Character{}, err
	}

	var a Character
	err = request.ToJSON(&a)

	return a, err
}

type Character struct {
	RequestHash        string          `json:"request_hash"`
	RequestCached      bool            `json:"request_cached"`
	RequestCacheExpiry int64           `json:"request_cache_expiry"`
	MalID              int64           `json:"mal_id"`
	URL                string          `json:"url"`
	Name               string          `json:"name"`
	NameKanji          string          `json:"name_kanji"`
	Nicknames          []string        `json:"nicknames"`
	About              string          `json:"about"`
	MemberFavorites    int64           `json:"member_favorites"`
	ImageURL           string          `json:"image_url"`
	Animeography       []common.Anime  `json:"animeography"`
	Mangaography       []common.Manga  `json:"mangaography"`
	VoiceActors        []common.Person `json:"voice_actors"`
}

func (r Character) GetID() interface{} {
	return r.MalID
}

func (r Character) GetType() string {
	return mal_types.Character
}
