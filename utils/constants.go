package utils

import (
    "fmt"
)

type constants struct {
    API string
}

var Contstants constants

func init() {
    Contstants = constants{
        API: "https://api.jikan.moe/v3",
    }
}

func (c *constants) AppendAPIf(endpoint string, values... interface{}) string {
    return fmt.Sprintf("%v%v", c.API, fmt.Sprintf(endpoint, values...))
}