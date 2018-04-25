package rss1

import (
	"encoding/json"

	"github.com/naoto0822/gss/module"
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
	Modules     module.Modules
}

// Image RSS1.0 image
type Image struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	URL   string `xml:"url"`
}

// Item RSS1.0 item
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Modules     module.Modules
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
	var date string
	if f.Channel.Modules.DublinCore.Date != "" {
		date = f.Channel.Modules.DublinCore.Date
	}

	gf := &struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
		PubDate     string `json:"pubdate"`
		Items       []Item `json:"items"`
	}{
		Title:       f.Channel.Title,
		Link:        f.Channel.Link,
		Description: f.Channel.Description,
		Image:       f.Image,
		PubDate:     date,
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
	var creator string
	if i.Modules.DublinCore.Creator != "" {
		creator = i.Modules.DublinCore.Creator
	}

	type author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	a := author{
		Name: creator,
	}
	var authors []author
	if a.Name != "" {
		authors = append(authors, a)
	}

	var date string
	if i.Modules.DublinCore.Date != "" {
		date = i.Modules.DublinCore.Date
	}

	gi := &struct {
		Title       string   `json:"title"`
		Link        string   `json:"link"`
		Description string   `json:"description"`
		PubDate     string   `json:"pubdate"`
		Authors     []author `json:"authors"`
	}{
		Title:       i.Title,
		Link:        i.Link,
		Description: i.Description,
		PubDate:     date,
		Authors:     authors,
	}
	return json.Marshal(gi)
}
