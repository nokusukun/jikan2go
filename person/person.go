// person implements Jikan's api/person resources and endpoints
package person

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetPerson(m common.MALItem) (Person, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/person/%v", m.GetID()))
	if err != nil {
		return Person{}, err
	}

	var a Person
	err = request.ToJSON(&a)

	return a, err
}

type Person struct {
	RequestHash         string               `json:"request_hash"`
	RequestCached       bool                 `json:"request_cached"`
	RequestCacheExpiry  int64                `json:"request_cache_expiry"`
	MalID               int64                `json:"mal_id"`
	URL                 string               `json:"url"`
	ImageURL            string               `json:"image_url"`
	WebsiteURL          string               `json:"website_url"`
	Name                string               `json:"name"`
	GivenName           string               `json:"given_name"`
	FamilyName          string               `json:"family_name"`
	AlternateNames      []string             `json:"alternate_names"`
	Birthday            string               `json:"birthday"`
	MemberFavorites     int64                `json:"member_favorites"`
	About               string               `json:"about"`
	VoiceActingRoles    []VoiceActingRole    `json:"voice_acting_roles"`
	AnimeStaffPositions []AnimeStaffPosition `json:"anime_staff_positions"`
	PublishedManga      []PublishedManga     `json:"published_manga"`
}

func (r Person) GetID() interface{} {
	return r.MalID
}

func (r Person) GetType() string {
	return mal_types.Person
}

type PublishedManga struct {
	Position string `json:"position"`
	Manga    Manga  `json:"manga"`
}

type AnimeStaffPosition struct {
	Position Position `json:"position"`
	Anime    Anime    `json:"anime"`
}

type Anime struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r Anime) GetID() interface{} {
	return r.MalID
}

func (r Anime) GetType() string {
	return mal_types.Anime
}

type Manga struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r Manga) GetID() interface{} {
	return r.MalID
}

func (r Manga) GetType() string {
	return mal_types.Manga
}

type Character struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r Character) GetID() interface{} {
	return r.MalID
}

func (r Character) GetType() string {
	return mal_types.Character
}

type VoiceActingRole struct {
	Role      Role      `json:"role"`
	Anime     Anime     `json:"anime"`
	Character Character `json:"character"`
}

type Position string

const (
	InsertedSongPerformance Position = "Inserted Song Performance"
	ThemeSongPerformance    Position = "Theme Song Performance"
)

type Role string

const (
	Main       Role = "Main"
	Supporting Role = "Supporting"
)
