package models

import "github.com/gorilla/mux"

type Router struct {
	Router *mux.Router
}

type CharacterResponse struct {
    Code           int             `json:"code"`
    Status         string          `json:"status"`
    Copyright      string          `json:"copyright"`
    AttributionText string          `json:"attributionText"`
    AttributionHTML string          `json:"attributionHTML"`
    Etag           string          `json:"etag"`
    Data           CharacterData   `json:"data"`
}

type CharacterData struct {
    Offset int           `json:"offset"`
    Limit  int           `json:"limit"`
    Total  int           `json:"total"`
    Count  int           `json:"count"`
    Results []CharacterResult `json:"results"`
}

type CharacterResult struct {
    ID          int            `json:"id"`
    Name        string         `json:"name"`
    Description string         `json:"description"`
    Modified    string         `json:"modified"`
    Thumbnail   CharacterThumbnail `json:"thumbnail"`
    ResourceURI string         `json:"resourceURI"`
    Comics      CharacterComics `json:"comics"`
    Series      CharacterSeries `json:"series"`
    Stories     CharacterStories `json:"stories"`
    Events      CharacterEvents `json:"events"`
    Urls        []CharacterURL `json:"urls"`
}

type CharacterThumbnail struct {
    Path      string `json:"path"`
    Extension string `json:"extension"`
}

// CharacterComics represents the comics details of a character
type CharacterComics struct {
    Available     int `json:"available"`
    CollectionURI string `json:"collectionURI"`
    Items         []CharacterComicItem `json:"items"`
    Returned      int `json:"returned"`
}


type CharacterComicItem struct {
    ResourceURI string `json:"resourceURI"`
    Name        string `json:"name"`
}


type CharacterSeries struct {
    Available     int `json:"available"`
    CollectionURI string `json:"collectionURI"`
    Items         []CharacterSeriesItem `json:"items"`
    Returned      int `json:"returned"`
}


type CharacterSeriesItem struct {
    ResourceURI string `json:"resourceURI"`
    Name        string `json:"name"`
}


type CharacterStories struct {
    Available     int `json:"available"`
    CollectionURI string `json:"collectionURI"`
    Items         []CharacterStoryItem `json:"items"`
    Returned      int `json:"returned"`
}


type CharacterStoryItem struct {
    ResourceURI string `json:"resourceURI"`
    Name        string `json:"name"`
    Type        string `json:"type"`
}


type CharacterEvents struct {
    Available     int `json:"available"`
    CollectionURI string `json:"collectionURI"`
    Items         []CharacterEventItem `json:"items"`
    Returned      int `json:"returned"`
}


type CharacterEventItem struct {
    ResourceURI string `json:"resourceURI"`
    Name        string `json:"name"`
}


type CharacterURL struct {
    Type string `json:"type"`
    Url  string `json:"url"`
}