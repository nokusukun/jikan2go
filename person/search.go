package person

type SearchResult struct {
	RequestHash        string   `json:"request_hash"`
	RequestCached      bool     `json:"request_cached"`
	RequestCacheExpiry int64    `json:"request_cache_expiry"`
	Results            []Result `json:"results"`
	LastPage           int64    `json:"last_page"`
}

type Result struct {
	MalID            int64    `json:"mal_id"`
	URL              string   `json:"url"`
	ImageURL         string   `json:"image_url"`
	Name             string   `json:"name"`
	AlternativeNames []string `json:"alternative_names"`
}

func (r Result) GetID() int64 {
	return r.MalID
}

func (r Result) GetType() string {
	return "person"
}
