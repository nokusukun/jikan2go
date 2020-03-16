package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_ByUsername(t *testing.T) {
	naux, err := GetUser(User{Username: "naux"})
	assert.Nil(t, err)

	assert.NotEqual(t, naux.Username, "")
	fmt.Println(naux.UserID, naux.Username)
	fmt.Println(naux.About)
}

func TestGetAnimeList(t *testing.T) {
	naux, err := GetUser(User{Username: "naux"})
	assert.Nil(t, err)

	_, err = GetUser(User{Username: "nekomata1037"})
	assert.Nil(t, err)

	_, err = GetUser(User{Username: "naux"})
	assert.Nil(t, err)

	nauxlist, err := GetAnimeList(naux, AllList)
	assert.Nil(t, err)

	assert.Greater(t, len(nauxlist.Anime), 0)
	fmt.Println(len(nauxlist.Anime))
	fmt.Println(nauxlist.Anime[0].Title)
}

func TestGetMangaList(t *testing.T) {
	nauxlist, err := GetMangaList(User{Username: "naux"}, AllList)
	assert.Nil(t, err)

	assert.Greater(t, len(nauxlist.Manga), 0)
	fmt.Println(len(nauxlist.Manga))
	fmt.Println(nauxlist.Manga[0].Title)
}

func TestGetFriends(t *testing.T) {
	nekoFriends, err := GetFriends(User{Username: "nekomata1037"}, 1)
	assert.Nil(t, err)

	assert.Greater(t, len(nekoFriends.Friends), 0)
	for _, i := range nekoFriends.Friends {
		fmt.Printf("%v ", i.Username)
	}
	fmt.Println("")
}
