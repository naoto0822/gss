package gss

import (
	"encoding/json"
	"html/template"
)

// Feed gss original feed
type Feed struct {
	RSSType     RSSType
	Title       string   `json:"title"`
	Links       []string `json:"links"`
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
	ID         string        `json:"id"`
	Title      string        `json:"title"`
	Links      []string      `json:"links"`
	Body       template.HTML `json:"body"`
	PubDate    string        `json:"pubdate"`
	Updated    string        `json:"updated"`
	Authors    []Author      `json:"authors"`
	Categories []string      `json:"categories"`
}

// Author gss author
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Map convert gss.Feed from other package Feed
func (f *Feed) Map(bytes []byte) error {
	return json.Unmarshal(bytes, f)
}
