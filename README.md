# jikan2go
--
A more go-like library for [Jikan](https://jikan.moe/).

### Download
```
$ go get github.com/nokusukun/jikan2go
```
Word of warning: This API doesn't do rate limiting yet.

## Usage

### Configuration
The endpoint used by the API can be changed by modifying `utils.Constants.API`
```go
package main

import (
    "fmt"

    "github.com/nokusukun/jikan2go/anime"
    "github.com/nokusukun/jikan2go/utils"
)

func init()  {
    utils.Contstants.API = "https://jikan.noku.pw/v3"   // All of jikan2go API calls will now use
                                                        // jikan.noku.pw instead of api.jikan.moe    
    
    utils.Constants.CacheDir = "~/.jikan/"  //change where cache files are stored (default: os.TempDir())
    utils.Constants.CacheLifetime = time.hour * 1 // change the cache lifetime before 
                                                  // requesting data from the remote resource. (default 30 minutes)                
        
}

```
***Note: Make sure to remove the trailing slash when specifying a new API endpoint.***
#### Note on caching
While the library maintains a cache, it also respects [Jikan's ETag headers.](https://jikan.docs.apiary.io/#introduction/cache-validation)

### Anime/Manga
#### Retrieve
#### Searching
Searching in jikan2go fully implements [Jikan's Search Parameters](https://jikan.docs.apiary.io/#reference/0/search/genre-request-example+schema?console=1)
```go
package main

import (
    "fmt"

    "github.com/nokusukun/jikan2go/anime"
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

    "github.com/nokusukun/jikan2go/anime"
    "github.com/nokusukun/jikan2go/common"
    "github.com/nokusukun/jikan2go/manga"
)

func main() {
    result, _ := anime.Search(anime.Query{Q:"made in abyss"})
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
```
Output
```
Title: Made in Abyss
Link: https://myanimelist.net/anime/34599/Made_in_Abyss
Members: 614368
Made in Abyss Author:  Tsukushi, Akihito

---
News related to Made in Abyss
Interview: Kinema Citrus Staff Reflect on 'Made in Abyss'
https://myanimelist.net/news/58244750

```

### Season
```go
package main

import (
    "fmt"

    "github.com/nokusukun/jikan2go/season"
)


func main() {

    // Print 5 anime shows that aired on summer 2019
    summer2019, err := season.GetSeason(season.Summer, 2019)
    if err != nil {
        panic(err)
    }

    for _, a := range summer2019.Anime[:5] {
        fmt.Println("Title:", a.Title)
        fmt.Println("Rating:", a.Score, "\n---")
    }
}
```
Output
```
Title: Dr. Stone
Rating: 8.46
---
Title: Enen no Shouboutai
Rating: 7.78
---
Title: Vinland Saga
Rating: 8.78
---
Title: Dungeon ni Deai wo Motomeru no wa Machigatteiru Darou ka II
Rating: 7.31
---
Title: Arifureta Shokugyou de Sekai Saikyou
Rating: 6.52
---

```

## Integration Status
* Package Testing
    * Working On
        * anime/anime
    * TBI
        * anime/character_staff
        * anime/episodes
        * anime/recommendations
        * anime/search
        * anime/videos
        * character/character
        * common/genre
        * common/news
        * common/pictures
        * common/reviews
        * common/stats
* Common
    * Implemented
        * review
        * genre
        * news
        * pictures
        * stats
* Anime
    * Implemented
        * **Request caching**
        * character_staff
        * episodes
        * recommendations
        * search
        * videos
    * Back Burner
        * forum
        * more info
        * user updates
* Manga
    * Implemented
        * characters
        * recommendations
        * search
    * Back Burner
        * forum
        * more info
        * user updates
* Person
    * Implemented
        * pictures
* Characters
    * Implemented
        * pictures
* Season
* Schedule
* Top
    * Anime
    * Manga
    * Person
    * Character
* Genre
* Producer
* Magazine

### To Be implemented
* ~~Request Caching~~
* User
* Club
* Meta