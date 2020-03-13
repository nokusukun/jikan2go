package main

import (
	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/character"
)

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// searching anime
	result, _ := anime.Search(anime.Query{Q: "made in abyss"})
	miaResult := result.Results[0]

	mia_id := miaResult.MalID
	madeInAbyss, err := anime.GetAnime(anime.Anime{MalID: mia_id})
	panicOn(err) // anime.GetAnime(miaResult) would also work as well

	miaStaff, err := anime.GetCharacterStaff(madeInAbyss)
	panicOn(err) // anime.GetCharacterStaff(anime.Anime{MalID:mia_id}) would also work as well

	nanachi, err := character.GetCharacter(miaStaff.Characters[0])
	panicOn(err)

}
