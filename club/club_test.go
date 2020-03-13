package club

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClub(t *testing.T) {
	club, err := GetClub(Club{MalID: 1})
	assert.Nil(t, err)
	assert.Contains(t, club.Title, "Cowboy")
	fmt.Println(club.Title)

	members, err := GetMembers(club, 1)
	assert.Nil(t, err)
	fmt.Println("Members:", len(members.Members))
}

func TestGetMembers(t *testing.T) {
	members, err := GetMembers(Club{MalID: 1}, 1)
	assert.Nil(t, err)
	assert.Greater(t, len(members.Members), 0)
	fmt.Println(len(members.Members))
}
