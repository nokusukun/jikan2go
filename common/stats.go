package common

import (
    "github.com/imroc/req"

    "github.com/nokusukun/Jikan2Go/utils"
)

func GetStats(m MALItem) (Stats, error) {
    request, err := req.Get(utils.Contstants.AppendAPIf("/%v/%v/stats",m.GetType(), m.GetID()))
    if err != nil {
        return Stats{}, err
    }

    var a Stats
    err = request.ToJSON(&a)

    return a, err
}

type Stats struct {
    RequestHash        string           `json:"request_hash"`
    RequestCached      bool             `json:"request_cached"`
    RequestCacheExpiry int64            `json:"request_cache_expiry"`
    Watching           int64            `json:"watching"`
    Completed          int64            `json:"completed"`
    OnHold             int64            `json:"on_hold"`
    Dropped            int64            `json:"dropped"`
    PlanToWatch        int64            `json:"plan_to_watch"`
    PlanToRead         int64            `json:"plan_to_read"`
    Total              int64            `json:"total"`
    Scores             map[string]Score `json:"scores"`
}

type Score struct {
    Votes      int64   `json:"votes"`
    Percentage float64 `json:"percentage"`
}

