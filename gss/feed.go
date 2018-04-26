package gss

import (
	"encoding/json"
)

// Feed gss original feed
type Feed struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Image       Image    `json:"image"`
	CopyRight   string   `json:"copyright"`
	PubDate     string   `json:"pubdate"`
	Updated     string   `json:"updated"`
	Authors     []Author `json:"authors"`
	Categories  []string `json:"categories"`
	Items       []Item   `json:"items"`
}

// Image gss image
type Image struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Link   string `json:"link"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Item gss item
type Item struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	PubDate     string    `json:"pubdate"`
	Updated     string    `json:"updated"`
	Authors     []Author  `json:"authors"`
	Categories  []string  `json:"categories"`
	Image       Image     `json:"image"`
	Enclosure   Enclosure `json:"enclosure"`
	// Thumbnail this is Media Module
	Thumbnail Thumbnail `json:"thumbnail"`
}

// Enclosure gss enclosure
type Enclosure struct {
	URL    string `json:"url"`
	Length int64  `json:"length"`
	Type   string `json:"type"`
}

// Author gss author
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Thumbnail gss thumbnail
// later move to Media Module type
type Thumbnail struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

// Map convert gss.Feed from other package Feed
func (f *Feed) Map(bytes []byte) error {
	return json.Unmarshal(bytes, f)
}
