package common

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

func GetForum(m MALItem) (Forum, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/%v/%v/forum", m.GetType(), m.GetID()))
	if err != nil {
		return Forum{}, err
	}

	var a Forum
	err = request.ToJSON(&a)

	return a, err
}

type Forum struct {
	RequestHash        string  `json:"request_hash"`
	RequestCached      bool    `json:"request_cached"`
	RequestCacheExpiry int64   `json:"request_cache_expiry"`
	Topics             []Topic `json:"topics"`
}

type Topic struct {
	TopicID    int64    `json:"topic_id"`
	URL        string   `json:"url"`
	Title      string   `json:"title"`
	DatePosted string   `json:"date_posted"`
	AuthorName string   `json:"author_name"`
	AuthorURL  string   `json:"author_url"`
	Replies    int64    `json:"replies"`
	LastPost   LastPost `json:"last_post"`
}

type LastPost struct {
	URL        string `json:"url"`
	AuthorName string `json:"author_name"`
	AuthorURL  string `json:"author_url"`
	DatePosted string `json:"date_posted"`
}
