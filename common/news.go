package common

import (
	"github.com/nokusukun/jikan2go/utils"
)

func GetNews(m MALItem) (News, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/%v/%v/news", m.GetType(), m.GetID()))
	if err != nil {
		return News{}, err
	}

	var a News
	err = request.ToJSON(&a)

	return a, err
}

type News struct {
	RequestHash        string    `json:"request_hash"`
	RequestCached      bool      `json:"request_cached"`
	RequestCacheExpiry int64     `json:"request_cache_expiry"`
	Articles           []Article `json:"articles"`
}

type Article struct {
	URL        string `json:"url"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	AuthorName string `json:"author_name"`
	AuthorURL  string `json:"author_url"`
	ForumURL   string `json:"forum_url"`
	ImageURL   string `json:"image_url"`
	Comments   int64  `json:"comments"`
	Intro      string `json:"intro"`
}
