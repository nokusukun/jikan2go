package main

import (
	"fmt"

	"github.com/nokusukun/jikan2go/anime"
	"github.com/nokusukun/jikan2go/common"
	"github.com/nokusukun/jikan2go/manga"
)

func main() {
	result, _ := anime.Search(anime.Query{Q: "made in abyss"})
	firstResult := result.Results[0]

	fmt.Println("Title:", firstResult.Title)
	fmt.Println("Link:", firstResult.URL)
	fmt.Println("Members:", firstResult.Members)

	madeInAbyss, _ := anime.GetAnime(firstResult)
	miaManga, _ := manga.GetManga(madeInAbyss.Related.Adaptation[0])
	fmt.Println("Made in Abyss Author: ", miaManga.Authors[0].Name)

	news, _ := common.GetNews(madeInAbyss)
	fmt.Println("\n---\nNews related to Made in Abyss")
	fmt.Println(news.Articles[0].Title)
	fmt.Println(news.Articles[0].URL)
}
