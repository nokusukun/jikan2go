package person

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

func GetPerson(m common.MALItem) (Person, error) {
	request, err := req.Get(utils.Constants.AppendAPIf("/person/%v", m.GetID()))
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
	WebsiteURL          interface{}          `json:"website_url"`
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

func (r Person) GetID() int64 {
	return r.MalID
}

func (r Person) GetType() string {
	return "person"
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

func (r Anime) GetID() int64 {
	return r.MalID
}

func (r Anime) GetType() string {
	return "anime"
}

type Manga struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r Manga) GetID() int64 {
	return r.MalID
}

func (r Manga) GetType() string {
	return "manga"
}

type Character struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r Character) GetID() int64 {
	return r.MalID
}

func (r Character) GetType() string {
	return "character"
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
