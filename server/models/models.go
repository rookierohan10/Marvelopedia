package models

import "github.com/gorilla/mux"

type Router struct {
	Router *mux.Router
}

type CharacterResponse struct {
	Code            int           `json:"code"`
	Status          string        `json:"status"`
	Copyright       string        `json:"copyright"`
	AttributionText string        `json:"attributionText"`
	AttributionHTML string        `json:"attributionHTML"`
	Etag            string        `json:"etag"`
	Data            CharacterData `json:"data"`
}

type CharacterSeriesResponse struct {
	Code            int        `json:"code"`
	Status          string     `json:"status"`
	Copyright       string     `json:"copyright"`
	AttributionText string     `json:"attributionText"`
	AttributionHTML string     `json:"attributionHTML"`
	Etag            string     `json:"etag"`
	Data            SeriesData `json:"data"`
}

type CharacterComicsResponse struct {
	Code            int        `json:"code"`
	Status          string     `json:"status"`
	Copyright       string     `json:"copyright"`
	AttributionText string     `json:"attributionText"`
	AttributionHTML string     `json:"attributionHTML"`
	Etag            string     `json:"etag"`
	Data            ComicsData `json:"data"`
}

type CharacterEventsResponse struct {
	Code            int        `json:"code"`
	Status          string     `json:"status"`
	Copyright       string     `json:"copyright"`
	AttributionText string     `json:"attributionText"`
	AttributionHTML string     `json:"attributionHTML"`
	Etag            string     `json:"etag"`
	Data            EventsData `json:"data"`
}

type CharacterStoriesResponse struct {
	Code            int         `json:"code"`
	Status          string      `json:"status"`
	Copyright       string      `json:"copyright"`
	AttributionText string      `json:"attributionText"`
	AttributionHTML string      `json:"attributionHTML"`
	Etag            string      `json:"etag"`
	Data            StoriesData `json:"data"`
}

type CharacterData struct {
	Offset  int               `json:"offset"`
	Limit   int               `json:"limit"`
	Total   int               `json:"total"`
	Count   int               `json:"count"`
	Results []CharacterResult `json:"results"`
}

type SeriesData struct {
	Offset  int            `json:"offset"`
	Limit   int            `json:"limit"`
	Total   int            `json:"total"`
	Count   int            `json:"count"`
	Results []SeriesResult `json:"results"`
}

type ComicsData struct {
	Offset  int            `json:"offset"`
	Limit   int            `json:"limit"`
	Total   int            `json:"total"`
	Count   int            `json:"count"`
	Results []ComicsResult `json:"results"`
}

type EventsData struct {
	Offset  int            `json:"offset"`
	Limit   int            `json:"limit"`
	Total   int            `json:"total"`
	Count   int            `json:"count"`
	Results []EventsResult `json:"results"`
}

type StoriesData struct {
	Offset  int             `json:"offset"`
	Limit   int             `json:"limit"`
	Total   int             `json:"total"`
	Count   int             `json:"count"`
	Results []StoriesResult `json:"results"`
}

type CharacterResult struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Modified    string             `json:"modified"`
	Thumbnail   CharacterThumbnail `json:"thumbnail"`
	ResourceURI string             `json:"resourceURI"`
	Comics      CharacterComics    `json:"comics"`
	Series      CharacterSeries    `json:"series"`
	Stories     CharacterStories   `json:"stories"`
	Events      CharacterEvents    `json:"events"`
	Urls        []CharacterURL     `json:"urls"`
}

type SeriesResult struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Urls        []CharacterURL     `json:"urls"`
	Startyear   int                `json:"startYear"`
	Endyear     int                `json:"endYear"`
	Rating      string             `json:"rating"`
	Type        string             `json:"type"`
	Modified    string             `json:"modified"`
	Thumbnail   CharacterThumbnail `json:"thumbnail"`
	Creators    Creators           `json:"creators"`
	Characters  Characters         `json:"characters"`
	Stories     CharacterStories   `json:"stories"`
	Comics      CharacterComics    `json:"comics"`
	Events      CharacterEvents    `json:"events"`
}

type ComicsResult struct {
	ID          int `json:"id`
	DigitalID   int
	Title       string
	IssueNumber int
	Description string
	Modified    string
	Format      string
	PageCount   int
	TextObjects []TextObjects
	Urls        []CharacterURL
	Series      CharacterSeriesItem
	Thumbnail   CharacterThumbnail
	Images []CharacterThumbnail
	Creators Creators
	Characters Characters
	Stories CharacterStories
	Events CharacterEvents
}

type EventsResult struct {
}

type StoriesResult struct {
}

type CharacterThumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type Creators struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []CreatorItem `json:"items"`
	Returned      int           `json:"returned"`
}

type CreatorItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

type Characters struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []CreatorItem `json:"items"`
	Returned      int           `json:"returned"`
}

// CharacterComics represents the comics details of a character
type CharacterComics struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterComicItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterComicItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type CharacterSeries struct {
	Available     int                   `json:"available"`
	CollectionURI string                `json:"collectionURI"`
	Items         []CharacterSeriesItem `json:"items"`
	Returned      int                   `json:"returned"`
}

type CharacterSeriesItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type CharacterStories struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterStoryItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterStoryItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type CharacterEvents struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterEventItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterEventItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type CharacterURL struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
