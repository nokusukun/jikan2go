package main

import (
	"fmt"

	"github.com/nokusukun/jikan2go/season"
)

func main() {

	// Print 5 anime shows that aired on summer 2019
	summer2019, err := season.GetSeason(season.Summer, 2019) // using 'summer' instead of season.Summer works as well
	if err != nil {
		panic(err)
	}

	for _, a := range summer2019.Anime[:5] {
		fmt.Println("Title:", a.Title)
		fmt.Println("Rating:", a.Score, "\n---")
	}
}
