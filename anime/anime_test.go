package anime

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAnime_GetID(t *testing.T) {
    miaID := int64(34599)

    MadeInAbyss, err := GetAnime(Anime{MalID: miaID})
    assert.Nil(t, err)
    fmt.Println("Title:", MadeInAbyss.Title)
    fmt.Println("Request Cached?", MadeInAbyss.RequestCached)
    assert.Equal(t, MadeInAbyss.MalID, miaID)
}

func TestSearch(t *testing.T) {
    result, err := Search(Query{Q: "made in abyss", Limit: 1, Genre: GenreAdventure})
    assert.Nil(t, err)
    assert.Len(t, result.Results, 1, result)
    fmt.Println("Title:", result.Results[0].Title)
    fmt.Println("Request Cached?", result.RequestCached)
}
