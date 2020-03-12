package manga

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	result, err := Search(Query{Q: "made in abyss"})
	assert.Nil(t, err)
	assert.Greater(t, len(result.Results), 0)

	manga := result.Results[0]
	fmt.Println("Manga ID:", manga.MalID)
	fmt.Println("Manga Name:", manga.Title)
	assert.Equal(t, manga.Title, "Made in Abyss")
}

func TestGetManga(t *testing.T) {
	mia, err := GetManga(Manga{MalID: 91941})
	assert.Nil(t, err)

	assert.Equal(t, mia.Title, "Made in Abyss")
	assert.Equal(t, mia.MalID, int64(91941))

	fmt.Println("Manga ID:", mia.MalID)
	fmt.Println("Manga Name:", mia.Title)
}

func TestGetCharacters(t *testing.T) {
	chars, err := GetCharacters(Manga{MalID: 91941})
	assert.Nil(t, err)
	assert.Greater(t, len(chars.Characters), 0)

	char := chars.Characters[0]

	fmt.Println("Character ID:", char.MalID)
	fmt.Println("Character Name:", char.Name)
}

func TestGetRecommendations(t *testing.T) {

	MadeInAbyss, err := GetManga(Manga{MalID: 91941})
	assert.Nil(t, err)

	recs, err := GetRecommendations(MadeInAbyss)
	assert.Nil(t, err)
	assert.Greater(t, len(recs.Recommendations), 0)
	fmt.Printf("Recommendation count: %v\n", len(recs.Recommendations))
	fmt.Printf("First Recommendation: %v\n", recs.Recommendations[0].Title)
}

func TestGetTop(t *testing.T) {
	result, err := GetTop(1, "")
	assert.Nil(t, err)
	assert.Greater(t, len(result.Top), 0)
	fmt.Println(result.Top[0].Title)

}

func TestGetTopSubtype(t *testing.T) {
	result, err := GetTop(1, SubTypeManga)
	assert.Nil(t, err)
	assert.Greater(t, len(result.Top), 0)
	fmt.Println(result.Top[0].Title)
}
