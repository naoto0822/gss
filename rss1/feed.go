package rss1

import (
	"encoding/json"
	"html/template"

	// implment interfaces.Mappable
	_ "github.com/naoto0822/gss/interfaces"
)

// cf. http://web.resource.org/rss/1.0/spec

// Feed RSS1.0 feed
type Feed struct {
	Channel   Channel   `xml:"channel"`
	Image     Image     `xml:"image"`
	Items     []Item    `xml:"item"`
	TextInput TextInput `xml:"textinput"`
}

// Channel RSS1.0 channel
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	// Date this is Dunblin Core
	Date string `xml:"date"`
	// Language this is Dunblin Core
	Language string `xml:"language"`
}

// Image RSS1.0 image
type Image struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	URL   string `xml:"url"`
}

// Item RSS1.0 item
type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Date this is Dunblin Core
	Date string `xml:"date"`
	// Creator this is Dunblin Core
	Creator string `xml:"creator"`
}

// TextInput RSS1.0 TextInput
type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

// ToJSON implemented interfaces.Mappable.
// convert to map gss.Feed.
func (f Feed) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}

// MarshalJSON assemble gss.Feed struct
func (f Feed) MarshalJSON() ([]byte, error) {
	var links []string
	if f.Channel.Link != "" {
		links = append(links, f.Channel.Link)
	}

	gf := &struct {
		Title       string   `json:"title"`
		Links       []string `json:"links"`
		Description string   `json:"description"`
		Image       Image    `json:"image"`
		PubDate     string   `json:"pubdate"`
		Items       []Item   `json:"items"`
	}{
		Title:       f.Channel.Title,
		Links:       links,
		Description: f.Channel.Description,
		Image:       f.Image,
		PubDate:     f.Channel.Date,
		Items:       f.Items,
	}
	return json.Marshal(gf)
}

// MarshalJSON assemble gss.Image struct
func (i Image) MarshalJSON() ([]byte, error) {
	gi := &struct {
		Title string `json:"title"`
		URL   string `json:"url"`
		Link  string `json:"link"`
	}{
		Title: i.Title,
		URL:   i.URL,
		Link:  i.Link,
	}
	return json.Marshal(gi)
}

// MarshalJSON assemble gss.Item struct
func (i Item) MarshalJSON() ([]byte, error) {
	var links []string
	if i.Link != "" {
		links = append(links, i.Link)
	}

	type author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	a := author{
		Name: i.Creator,
	}
	var authors []author
	if a.Name != "" {
		authors = append(authors, a)
	}

	gi := &struct {
		Title       string        `json:"title"`
		Links       []string      `json:"links"`
		Description template.HTML `json:"description"`
		PubDate     string        `json:"pubdate"`
		Authors     []author      `json:"authors"`
	}{
		Title:       i.Title,
		Links:       links,
		Description: i.Description,
		PubDate:     i.Date,
		Authors:     authors,
	}
	return json.Marshal(gi)
}
