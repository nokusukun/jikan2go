package utils

import (
	"fmt"
	"os"
	"time"
)

type constants struct {
	API           string
	CacheDir      string
	CacheLifetime time.Duration
}

var Constants constants

func init() {
	Constants = constants{
		API:           "https://api.jikan.moe/v3",
		CacheDir:      os.TempDir(),
		CacheLifetime: time.Minute * 30,
	}
}

func (c *constants) AppendAPIf(endpoint string, values ...interface{}) string {
	return fmt.Sprintf("%v%v", c.API, fmt.Sprintf(endpoint, values...))
}
