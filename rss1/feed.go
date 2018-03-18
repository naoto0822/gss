package rss1

import (
	"encoding/json"
	"html/template"

	"github.com/naoto0822/gss/interfaces"
)

// cf. http://web.resource.org/rss/1.0/spec

// Feed RSS1.0 feed
type Feed struct {
	interfaces.Mappable
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
	Date        string `xml:"date"`
	Language    string `xml:"language"`
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
	Date        string        `xml:"date"`
	Creator     string        `xml:"creator"`
}

// TextInput RSS1.0 TextInput
type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

func (f *Feed) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}
