package user

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

type Filter string

/*
watching	/reading
/completed	/completed
/onhold	/onhold
/dropped	/dropped
/plantowatch	/plantoread
/ptw (alias)
*/
const (
	AllList         Filter = "all"
	WatchingList    Filter = "watching"
	ReadingList     Filter = "reading"
	CompletedList   Filter = "completed"
	OnHoldList      Filter = "onhold"
	PlanToWatchList Filter = "plantowatch"
	PlanToReadList  Filter = "plantoread"
)

func GetAnimeList(user common.MALItem, filter Filter) (AnimeList, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/user/%v/animelist/%v", user.GetID(), filter))
	if err != nil {
		return AnimeList{}, err
	}

	var a AnimeList
	err = request.ToJSON(&a)

	return a, err
}

func GetMangaList(user common.MALItem, filter Filter) (MangaList, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/user/%v/mangalist/%v", user.GetID(), filter))
	if err != nil {
		return MangaList{}, err
	}

	var a MangaList
	err = request.ToJSON(&a)

	return a, err
}

type AnimeList struct {
	RequestHash        string          `json:"request_hash"`
	RequestCached      bool            `json:"request_cached"`
	RequestCacheExpiry int64           `json:"request_cache_expiry"`
	Anime              []AnimeListItem `json:"anime"`
}

type MangaList struct {
	RequestHash        string          `json:"request_hash"`
	RequestCached      bool            `json:"request_cached"`
	RequestCacheExpiry int64           `json:"request_cache_expiry"`
	Manga              []MangaListItem `json:"manga"`
}

type AnimeListItem struct {
	MalID           int64       `json:"mal_id"`
	Title           string      `json:"title"`
	VideoURL        string      `json:"video_url"`
	URL             string      `json:"url"`
	ImageURL        string      `json:"image_url"`
	Type            string      `json:"type"`
	WatchingStatus  int64       `json:"watching_status"`
	Score           int64       `json:"score"`
	WatchedEpisodes int64       `json:"watched_episodes"`
	TotalEpisodes   int64       `json:"total_episodes"`
	AiringStatus    int64       `json:"airing_status"`
	SeasonName      interface{} `json:"season_name"`
	SeasonYear      interface{} `json:"season_year"`
	HasEpisodeVideo bool        `json:"has_episode_video"`
	HasPromoVideo   bool        `json:"has_promo_video"`
	HasVideo        bool        `json:"has_video"`
	IsRewatching    bool        `json:"is_rewatching"`
	Tags            interface{} `json:"tags"`
	Rating          string      `json:"rating"`
	StartDate       string      `json:"start_date"`
	EndDate         string      `json:"end_date"`
	WatchStartDate  interface{} `json:"watch_start_date"`
	WatchEndDate    interface{} `json:"watch_end_date"`
	Days            interface{} `json:"days"`
	Storage         interface{} `json:"storage"`
	Priority        string      `json:"priority"`
	AddedToList     bool        `json:"added_to_list"`
	Studios         []string    `json:"studios"`
	Licensors       []string    `json:"licensors"`
}

type MangaListItem struct {
	MalID            int64       `json:"mal_id"`
	Title            string      `json:"title"`
	URL              string      `json:"url"`
	ImageURL         string      `json:"image_url"`
	Type             string      `json:"type"`
	ReadingStatus    int64       `json:"reading_status"`
	Score            int64       `json:"score"`
	ReadChapters     int64       `json:"read_chapters"`
	ReadVolumes      int64       `json:"read_volumes"`
	TotalChapters    int64       `json:"total_chapters"`
	TotalVolumes     int64       `json:"total_volumes"`
	PublishingStatus int64       `json:"publishing_status"`
	IsRereading      bool        `json:"is_rereading"`
	Tags             interface{} `json:"tags"`
	StartDate        string      `json:"start_date"`
	EndDate          string      `json:"end_date"`
	ReadStartDate    interface{} `json:"read_start_date"`
	ReadEndDate      interface{} `json:"read_end_date"`
	Days             interface{} `json:"days"`
	Retail           interface{} `json:"retail"`
	Priority         string      `json:"priority"`
	AddedToList      bool        `json:"added_to_list"`
	Magazines        []string    `json:"magazines"`
}
