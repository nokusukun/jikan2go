package main

import (
	"fmt"

	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/utils"
)

func init() {
	utils.Contstants.API = "https://jikan.noku.pw/v3" // All of jikan2go API calls will now use
	// jikan.noku.pw instead of api.jikan.moe
}

func main() {
	utils.Contstants.API = "https://jikan.noku.pw/v3" // All of jikan2go API calls will now use
	// jikan.noku.pw instead of api.jikan.moe

	pet, err := anime.GetAnime(anime.Anime{MalID: 34599})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pet Description\n\n%v", pet.Synopsis)
	fmt.Println("Pet Rating", pet.Rating)
}
