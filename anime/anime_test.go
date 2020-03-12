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
	result, err := Search(
		Query{Q: "made in abyss",
			Limit: 1,
			Genre: GenreAdventure,
		})
	assert.Nil(t, err)

	assert.Len(t, result.Results, 1, result)

	fmt.Println("Title:", result.Results[0].Title)
	fmt.Println("Request Cached?", result.RequestCached)
}

func TestGetCharacterStaff(t *testing.T) {
	miaID := int64(34599)
	MadeInAbyss, err := GetAnime(Anime{MalID: miaID})
	assert.Nil(t, err)

	staff, err := GetCharacterStaff(MadeInAbyss)
	assert.Nil(t, err)
	assert.NotEqual(t, len(staff.Characters), 0)
	fmt.Println(staff.Characters[0].Name)
}

func TestGetEpisodes(t *testing.T) {
	miaID := int64(34599)

	MadeInAbyss, err := GetAnime(Anime{MalID: miaID})
	assert.Nil(t, err)

	episodes, err := GetEpisodes(MadeInAbyss)
	assert.Nil(t, err)
	assert.Equal(t, len(episodes.Episodes), 13)
	fmt.Printf("%+v\n", episodes.Episodes[0])
	assert.Equal(t, episodes.Episodes[0].Title, "The City of the Great Pit")
}

func TestGetRecommendations(t *testing.T) {
	miaID := int64(34599)

	MadeInAbyss, err := GetAnime(Anime{MalID: miaID})
	assert.Nil(t, err)

	recs, err := GetRecommendations(MadeInAbyss)
	assert.Nil(t, err)
	assert.Greater(t, len(recs.Recommendations), 12)
	fmt.Printf("Recommendations: %v\n", len(recs.Recommendations))
}

func TestGetTop(t *testing.T) {
	result, err := GetTop(1, "")
	assert.Nil(t, err)
	assert.Greater(t, len(result.Top), 0)
	fmt.Println(result.Top[0].Title)

}

func TestGetTopSubtype(t *testing.T) {
	result, err := GetTop(1, SubTypeMovie)
	assert.Nil(t, err)
	assert.Greater(t, len(result.Top), 0)
	fmt.Println(result.Top[0].Title)
}

func TestFMAB(t *testing.T) {
	result, err := Search(Query{Q: "FMAB", Order: OrderScore})
	assert.Nil(t, err)
	assert.Greater(t, len(result.Results), 0)

	fma := result.Results[0]
	fmt.Println("Title:", fma.Title)
	fmt.Println("Members: ", fma.Members)
}
