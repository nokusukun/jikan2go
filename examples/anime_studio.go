package main

import (
	"fmt"

	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/season"
	"github.com/nokusukun/jikan2go/studio"
)

func main() {

	// Print 5 anime shows that aired on summer 2019
	summer2019, err := season.GetSeason(season.Summer, 2019) // using 'summer' instead of season.Summer works as well
	if err != nil {
		panic(err)
	}

	for _, a := range summer2019.Anime[:5] {
		fmt.Println("Title:", a.Title)
		fmt.Println("Rating:", a.Score)

		// Since studio details are not available on GetSeason results, we'll retrieve the full anime object.
		fullAnime, _ := anime.GetAnime(a)

		s, err := studio.GetStudio(fullAnime.Studios[0], 1)
		if err != nil {
			fmt.Println("Error retrieving studio: ", err)
		} else {
			fmt.Println("Other anime by", s.Meta.Name, ":", s.Anime[0].Title)
		}

		fmt.Println("---")
	}
}
