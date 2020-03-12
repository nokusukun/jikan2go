package season

import (
	"github.com/imroc/req"

	"github.com/nokusukun/jikan2go/utils"
)

type Day string

const (
	Monday      Day = "monday"
	Tuesday     Day = "tuesday"
	Wednesday   Day = "wednesday"
	Thursday    Day = "thursday"
	Friday      Day = "friday"
	Saturday    Day = "saturday"
	Sunday      Day = "sunday"
	AllSchedule Day = ""
)

func GetSchedule(day Day) (Schedule, error) {
	request, err := req.Get(utils.Config.AppendAPIf("/schedule/%v", day))
	if err != nil {
		return Schedule{}, err
	}

	var a Schedule
	err = request.ToJSON(&a)

	return a, err
}

type Schedule struct {
	RequestHash        string         `json:"request_hash"`
	RequestCached      bool           `json:"request_cached"`
	RequestCacheExpiry int64          `json:"request_cache_expiry"`
	Monday             []AnimeElement `json:"monday"`
	Tuesday            []AnimeElement `json:"tuesday"`
	Wednesday          []AnimeElement `json:"wednesday"`
	Thursday           []AnimeElement `json:"thursday"`
	Friday             []AnimeElement `json:"friday"`
	Saturday           []AnimeElement `json:"saturday"`
	Sunday             []AnimeElement `json:"sunday"`
	Other              []AnimeElement `json:"other"`
	Unknown            []AnimeElement `json:"unknown"`
}
