package main

import (
    "fmt"

    "github.com/nokusukun/Jikan2Go/anime"
    "github.com/nokusukun/Jikan2Go/utils"
)

func init()  {
    utils.Contstants.API = "https://jikan.noku.pw/v3"   // All of Jikan2Go API calls will now use
                                                        // jikan.noku.pw instead of api.jikan.moe
}

func main() {
    utils.Contstants.API = "https://jikan.noku.pw/v3"   // All of Jikan2Go API calls will now use
                                                        // jikan.noku.pw instead of api.jikan.moe

    pet, err := anime.GetAnime(anime.Anime{MalID:34599})
    if err != nil {
        panic(err)
    }

    fmt.Printf("Pet Description\n\n%v", pet.Synopsis)
    fmt.Println("Pet Rating", pet.Rating)
}
