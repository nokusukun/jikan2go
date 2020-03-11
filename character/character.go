package character

import (
    "github.com/imroc/req"

    "github.com/nokusukun/Jikan2Go/common"
    "github.com/nokusukun/Jikan2Go/utils"
)

func GetCharacter(m common.MALItem) (Character, error) {
    request, err := req.Get(utils.Contstants.AppendAPIf("/character/%v", m.GetID()))
    if err != nil {
        return Character{}, err
    }

    var a Character
    err = request.ToJSON(&a)

    return a, err
}

type Character struct {
    RequestHash        string         `json:"request_hash"`
    RequestCached      bool           `json:"request_cached"`
    RequestCacheExpiry int64          `json:"request_cache_expiry"`
    MalID              int64          `json:"mal_id"`
    URL                string         `json:"url"`
    Name               string         `json:"name"`
    NameKanji          string         `json:"name_kanji"`
    Nicknames          []interface{}  `json:"nicknames"`
    About              string         `json:"about"`
    MemberFavorites    int64          `json:"member_favorites"`
    ImageURL           string         `json:"image_url"`
    Animeography       []interface{}  `json:"animeography"`
    Mangaography       []Mangaography `json:"mangaography"`
    VoiceActors        []interface{}  `json:"voice_actors"`
}

func (r Character) GetID() int64 {
    return r.MalID
}

func (r Character) GetType() string {
    return "character"
}

type Mangaography struct {
    MalID    int64  `json:"mal_id"`
    Name     string `json:"name"`
    URL      string `json:"url"`
    ImageURL string `json:"image_url"`
    Role     string `json:"role"`
}

