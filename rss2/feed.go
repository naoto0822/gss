package rss2

import (
	"encoding/json"

	// implement interfaces.Mappable
	_ "github.com/naoto0822/gss/interfaces"
)

// Feed RSS2.0 feed
type Feed struct {
	Channel Channel `xml:"channel"`
}

// Channel RSS2.0 channel elements
type Channel struct {
	Title          string     `xml:"title"`
	Link           string     `xml:"link"`
	Description    string     `xml:"description"`
	Language       string     `xml:"language"`
	CopyRight      string     `xml:"copyright"`
	ManagingEditor string     `xml:"managingEditor"`
	WebMaster      string     `xml:"webMaster"`
	PubDate        string     `xml:"pubDate"`
	LastBuildDate  string     `xml:"lastBuildDate"`
	Categories     []Category `xml:"category"`
	Generator      string     `xml:"generator"`
	Docs           string     `xml:"docs"`
	Cloud          Cloud      `xml:"cloud"`
	TTL            int        `xml:"ttl"`
	Image          Image      `xml:"image"`
	Rating         string     `xml:"rating"`
	TextInput      TextInput  `xml:"textInput"`
	SkipHours      SkipHours  `xml:"skipHours"`
	SkipDays       SkipDays   `xml:"skipDays"`
	Items          []Item     `xml:"item"`
}

// Category RSS2.0 category elements
type Category struct {
	Value  string `xml:",chardata"`
	Domain string `xml:"domain,attr"`
}

// Cloud RSS2.0 cloud elements
type Cloud struct {
	Domain            string `xml:"domain,attr"`
	Port              string `xml:"port,attr"`
	Path              string `xml:"path,attr"`
	RegisterProcedure string `xml:"registerProcedure,attr"`
	Protocol          string `xml:"protocol,attr"`
}

// Image RSS2.0 image elements
type Image struct {
	URL         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Width       int    `xml:"width"`
	Height      int    `xml:"height"`
}

// TextInput RSS2.0 testInput elements
type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

// SkipHours RSS2.0 skipHours elements
type SkipHours struct {
	Hours []int `xml:"hour"`
}

// SkipDays RSS2.0 skipDays elements
type SkipDays struct {
	Days []string `xml:"day"`
}

// Item RSS2.0 item elements
type Item struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Author      string     `xml:"author"`
	Categories  []Category `xml:"category"`
	Comments    string     `xml:"comments"`
	Enclosure   Enclosure  `xml:"enclosure"`
	GUID        GUID       `xml:"guid"`
	PubDate     string     `xml:"pubDate"`
	Source      Source     `xml:"source"`
	// Content this is Content Module
	Content string `xml:"encoded"`
	// Thumbnail this is Media Module
	Thumbnail Thumbnail `xml:"thumbnail"`
}

// Enclosure RSS2.0 enclosure elements
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// GUID RSS2.0 guid elements
type GUID struct {
	Value       string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

// Source RSS2.0 source elements
type Source struct {
	Value string `xml:",chardata"`
	URL   string `xml:"url,attr"`
}

// Thumbnail Media Module
type Thumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int64  `xml:"width,attr"`
	Height int64  `xml:"height,attr"`
}

// ToJSON implemented interfaces.Mappable
// convert to gss.Feed
func (f Feed) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}

// MarshalJSON assemble gss.Feed struct
func (f Feed) MarshalJSON() ([]byte, error) {
	var categories []Category
	if len(f.Channel.Categories) > 0 {
		categories = f.Channel.Categories
	}

	gf := &struct {
		Title       string     `json:"title"`
		Link        string     `json:"link"`
		Description string     `json:"description"`
		Image       Image      `json:"image"`
		CopyRight   string     `json:"copyright"`
		PubDate     string     `json:"pubdate"`
		Updated     string     `json:"updated"`
		Categories  []Category `json:"categories"`
		Items       []Item     `json:"items"`
	}{
		Title:       f.Channel.Title,
		Link:        f.Channel.Link,
		Description: f.Channel.Description,
		Image:       f.Channel.Image,
		CopyRight:   f.Channel.CopyRight,
		PubDate:     f.Channel.PubDate,
		Updated:     f.Channel.LastBuildDate,
		Categories:  categories,
		Items:       f.Channel.Items,
	}
	return json.Marshal(gf)
}

// MarshalJSON assemble gss.Image struct
func (i Image) MarshalJSON() ([]byte, error) {
	gi := &struct {
		Title  string `json:"title"`
		URL    string `json:"url"`
		Link   string `json:"link"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	}{
		Title:  i.Title,
		URL:    i.URL,
		Link:   i.Link,
		Width:  i.Width,
		Height: i.Height,
	}
	return json.Marshal(gi)
}

// MarshalJSON assemble gss.Category struct
func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Value)
}

// MarshalJSON assemble gss.Enclosure struct
func (e Enclosure) MarshalJSON() ([]byte, error) {
	ge := &struct {
		URL    string `json:"url"`
		Length int64  `json:"length"`
		Type   string `json:"type"`
	}{
		URL:    e.URL,
		Length: e.Length,
		Type:   e.Type,
	}
	return json.Marshal(ge)
}

// MarshalJSON assemble gss.Item struct
func (i Item) MarshalJSON() ([]byte, error) {
	type author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	a := author{
		Name: i.Author,
	}
	var authors []author
	if a.Name != "" {
		authors = append(authors, a)
	}

	gi := &struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Link        string     `json:"link"`
		Description string     `json:"description"`
		Content     string     `json:"content"`
		PubDate     string     `json:"pubdate"`
		Authors     []author   `json:"authors"`
		Categories  []Category `json:"categories"`
		Enclosure   Enclosure  `json:"enclosure"`
		Thumbnail   Thumbnail  `json:"thumbnail"`
	}{
		ID:          i.GUID.Value,
		Title:       i.Title,
		Link:        i.Link,
		Description: i.Description,
		Content:     i.Content,
		PubDate:     i.PubDate,
		Authors:     authors,
		Categories:  i.Categories,
		Enclosure:   i.Enclosure,
		Thumbnail:   i.Thumbnail,
	}
	return json.Marshal(gi)
}

// MarshalJSON assemble gss.Thumbnail struct
func (t Thumbnail) MarshalJSON() ([]byte, error) {
	gt := &struct {
		URL    string `json:"url"`
		Width  int64  `json:"width"`
		Height int64  `json:"height"`
	}{
		URL:    t.URL,
		Width:  t.Width,
		Height: t.Height,
	}
	return json.Marshal(gt)
}
