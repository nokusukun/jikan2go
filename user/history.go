package user

import (
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/utils"
)

type HistoryType = string

const (
	HistoryAll   HistoryType = ""
	HistoryManga HistoryType = "manga"
	HistoryAnime HistoryType = "anime"
)

func GetHistory(m common.MALItem, history_type HistoryType) (History, error) {
	request, err := utils.CachedReqGet(utils.Config.AppendAPIf("/user/%v/history/%v", m.GetID(), history_type))
	if err != nil {
		return History{}, err
	}

	var a History
	err = request.ToJSON(&a)

	return a, err
}

type History struct {
	RequestHash        string           `json:"request_hash"`
	RequestCached      bool             `json:"request_cached"`
	RequestCacheExpiry int64            `json:"request_cache_expiry"`
	History            []HistoryElement `json:"history"`
}

type HistoryElement struct {
	Meta      common.TypedMALItem `json:"meta"`
	Increment int64               `json:"increment"`
	Date      string              `json:"date"`
}
