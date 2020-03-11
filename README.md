# Jikan2Go
--
A more go-like library for [Jikan](https://jikan.moe/).

### Download
```
$ go get github.com/nokusukun/Jikan2Go
```

## Usage

### Configuration
The endpoint used by the API can be changed by modifying `utils.Constants.API`
```go
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
```
***Note: Make sure to remove the trailing slash when specifying a new API endpoint.***

### Anime/Manga

#### Searching
Searching in Jikan2Go fully implements [Jikan's Search Parameters](https://jikan.docs.apiary.io/#reference/0/search/genre-request-example+schema?console=1)
```go
package main

import (
    "fmt"

    "github.com/nokusukun/Jikan2Go/anime"
)

func main() {
    // Searching for Anime
    result, _ := anime.Search(anime.Query{Q:"made in abyss"}) // same goes for manga.Search

    madeInAbyss := result.Results[0]

    fmt.Println("Title:", madeInAbyss.Title)
    fmt.Println("Link:", madeInAbyss.URL)
    
    // Retrieve recommendations for this particular anime
    recs, _ := anime.GetRecommendations(madeInAbyss)
    for _, recAnime := range recs.Recommendations[:3] {
        fmt.Println(recAnime.Title, "is recommended by", recAnime.RecommendationCount, "users")
    }

    fmt.Println("---")
    // Retrieving anime by it's MAL ID
    mia, err := anime.GetAnime(anime.Anime{MalID:34599})
    if err != nil {
        panic(err)
    }

    fmt.Printf("Made in Abyss Description\n\n%v", mia.Synopsis)
}
```
Output
```
Title: Made in Abyss
Link: https://myanimelist.net/anime/34599/Made_in_Abyss
Shinsekai yori is recommended by 22 users
Hunter x Hunter (2011) is recommended by 21 users
Yakusoku no Neverland is recommended by 21 users
---
Made in Abyss Description

The Abyss—a gaping chasm stretching down into the depths of the earth, filled with mysterious creatures and relics from 
a time long past. How did it come to be? What lies at the bottom? Countless brave individuals, known as Divers, have 
sought to solve these mysteries of the Abyss, fearlessly descending into its darkest realms.
```

***Note: The struct returned by anime.Search does not include the full canonical data, in this case, just feed the
    search result to anime.GetAnime***
```go
func main() {
    result, _ := anime.Search(anime.Query{Q:"made in abyss"})
    firstResult := result.Results[0]

    fmt.Println("Title:", firstResult.Title)
    fmt.Println("Link:", firstResult.URL)

    madeInAbyss, _ := anime.GetAnime(firstResult)
    miaManga, _ := manga.GetManga(madeInAbyss.Related.Adaptation[0])
    fmt.Println("Made in Abyss Author: ", miaManga.Authors[0].Name)
}
```
Output
```
Title: Made in Abyss
Link: https://myanimelist.net/anime/34599/Made_in_Abyss
Made in Abyss Author:  Tsukushi, Akihito

```

### Common Elements
The api also implements several common elements such as `genre`, `news`, `pictures`, `reviews` and `stats`.

```go
package main

import (
    "fmt"

    "github.com/nokusukun/Jikan2Go/anime"
    "github.com/nokusukun/Jikan2Go/common"
    "github.com/nokusukun/Jikan2Go/manga"
)

func main() {
    result, _ := anime.Search(anime.Query{Q:"made in abyss"})
    firstResult := result.Results[0]

    fmt.Println("Title:", firstResult.Title)
    fmt.Println("Link:", firstResult.URL)


    madeInAbyss, _ := anime.GetAnime(firstResult)
    miaManga, _ := manga.GetManga(madeInAbyss.Related.Adaptation[0])
    fmt.Println("Made in Abyss Author: ", miaManga.Authors[0].Name)

    news, _ := common.GetNews(madeInAbyss)
    fmt.Println("\n---\nNews related to Made in Abyss")
    fmt.Println(news.Articles[0].Title)
    fmt.Println(news.Articles[0].URL)
}
```
Output
```
Title: Made in Abyss
Link: https://myanimelist.net/anime/34599/Made_in_Abyss
Made in Abyss Author:  Tsukushi, Akihito

---
News related to Made in Abyss
Interview: Kinema Citrus Staff Reflect on 'Made in Abyss'
https://myanimelist.net/news/58244750

```

## To implement
* Tests
* Top
* Genre
* Producer
* Magazine
* User
* Club
* Meta