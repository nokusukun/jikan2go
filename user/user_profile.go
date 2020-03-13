package user

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/mal_types"
	"github.com/nokusukun/jikan2go/utils"
)

func GetUser(user common.MALItem) (User, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/user/%v", user.GetID()))
	if err != nil {
		return User{}, err
	}

	var a User
	err = request.ToJSON(&a)

	return a, err
}

type User struct {
	RequestHash        string     `json:"request_hash"`
	RequestCached      bool       `json:"request_cached"`
	RequestCacheExpiry int64      `json:"request_cache_expiry"`
	UserID             int64      `json:"user_id"`
	Username           string     `json:"username"`
	URL                string     `json:"url"`
	ImageURL           string     `json:"image_url"`
	LastOnline         string     `json:"last_online"`
	Gender             string     `json:"gender"`
	Birthday           string     `json:"birthday"`
	Location           string     `json:"location"`
	Joined             string     `json:"joined"`
	AnimeStats         AnimeStats `json:"anime_stats"`
	MangaStats         MangaStats `json:"manga_stats"`
	Favorites          Favorites  `json:"favorites"`
	About              string     `json:"about"`
}

func (r User) GetID() interface{} {
	return r.Username
}

func (r User) GetType() string {
	return mal_types.User
}

type AnimeStats struct {
	DaysWatched     float64 `json:"days_watched"`
	MeanScore       float64 `json:"mean_score"`
	Watching        int64   `json:"watching"`
	Completed       int64   `json:"completed"`
	OnHold          int64   `json:"on_hold"`
	Dropped         int64   `json:"dropped"`
	PlanToWatch     int64   `json:"plan_to_watch"`
	TotalEntries    int64   `json:"total_entries"`
	Rewatched       int64   `json:"rewatched"`
	EpisodesWatched int64   `json:"episodes_watched"`
}

type MangaStats struct {
	DaysRead     float64 `json:"days_read"`
	MeanScore    float64 `json:"mean_score"`
	Reading      int64   `json:"reading"`
	Completed    int64   `json:"completed"`
	OnHold       int64   `json:"on_hold"`
	Dropped      int64   `json:"dropped"`
	PlanToRead   int64   `json:"plan_to_read"`
	TotalEntries int64   `json:"total_entries"`
	Reread       int64   `json:"reread"`
	ChaptersRead int64   `json:"chapters_read"`
	VolumesRead  int64   `json:"volumes_read"`
}

type Favorites struct {
	Anime      []FavoriteAnime     `json:"anime"`
	Manga      []FavoriteManga     `json:"manga"`
	Characters []FavoriteCharacter `json:"characters"`
	People     []FavoritePerson    `json:"people"`
}

type FavoriteAnime struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r FavoriteAnime) GetID() interface{} {
	return r.MalID
}

func (r FavoriteAnime) GetType() string {
	return mal_types.Anime
}

type FavoriteManga struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r FavoriteManga) GetID() interface{} {
	return r.MalID
}

func (r FavoriteManga) GetType() string {
	return mal_types.Manga
}

type FavoriteCharacter struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r FavoriteCharacter) GetID() interface{} {
	return r.MalID
}

func (r FavoriteCharacter) GetType() string {
	return mal_types.Character
}

type FavoritePerson struct {
	MalID    int64  `json:"mal_id"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func (r FavoritePerson) GetID() interface{} {
	return r.MalID
}

func (r FavoritePerson) GetType() string {
	return mal_types.Person
}
