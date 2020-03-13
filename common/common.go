// common includes several misc methods and structs that are shared between different resources and endpoints
package common

import (
	"github.com/nokusukun/jikan2go/mal_types"
)

type MALItem interface {
	GetID() interface{}
	GetType() string
}

type TypedMALItem struct {
	MalID int64  `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (r TypedMALItem) GetID() interface{} {
	return r.MalID
}

func (r TypedMALItem) GetType() string {
	return r.Type
}

type Anime struct {
	MalID    int64  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (r Anime) GetID() interface{} {
	return r.MalID
}

func (r Anime) GetType() string {
	return mal_types.Anime
}

type Manga struct {
	MalID    int64  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (r Manga) GetID() interface{} {
	return r.MalID
}

func (r Manga) GetType() string {
	return mal_types.Manga
}

type Character struct {
	MalID    int64  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (r Character) GetID() interface{} {
	return r.MalID
}

func (r Character) GetType() string {
	return mal_types.Character
}

type Person struct {
	MalID    int64  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Language string `json:"language"`
}

func (r Person) GetID() interface{} {
	return r.MalID
}

func (r Person) GetType() string {
	return mal_types.Person
}

type User struct {
	MalID    int64  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (r User) GetID() interface{} {
	return r.Name
}

func (r User) GetType() string {
	return mal_types.User
}

type Member struct {
	Username string `json:"username"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (r Member) GetID() interface{} {
	return r.Username
}

func (r Member) GetType() string {
	return mal_types.User
}
