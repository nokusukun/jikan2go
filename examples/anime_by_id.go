package main

import (
	"fmt"

	"github.com/nokusukun/jikan2go/anime"
)

func main() {

	pet, err := anime.GetAnime(anime.Anime{MalID: 34599})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pet Description\n\n%v", pet.Synopsis)
	fmt.Println("Pet Rating", pet.Rating)
}
