package common

type Genre struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r Genre) GetID() interface{} {
	return r.MalID
}

func (r Genre) GetType() string {
	return "genre"
}
