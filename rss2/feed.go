package rss2

import (
	"encoding/json"

	"github.com/naoto0822/gss/modules"
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
	Modules        modules.Modules
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
	Modules     modules.Modules
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

	var date string
	if f.Channel.PubDate != "" {
		date = f.Channel.PubDate
	} else if f.Channel.Modules.DublinCore.Date != "" {
		date = f.Channel.Modules.DublinCore.Date
	}

	var updated string
	if f.Channel.LastBuildDate != "" {
		updated = f.Channel.LastBuildDate
	} else if f.Channel.Modules.DublinCore.Modified != "" {
		updated = f.Channel.Modules.DublinCore.Modified
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
		PubDate:     date,
		Updated:     updated,
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
	var authorName string
	if i.Author != "" {
		authorName = i.Author
	} else if i.Modules.DublinCore.Creator != "" {
		authorName = i.Modules.DublinCore.Creator
	}

	type author struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	a := author{
		Name: authorName,
	}
	var authors []author
	if a.Name != "" {
		authors = append(authors, a)
	}

	var pubDate string
	if i.PubDate != "" {
		pubDate = i.PubDate
	} else if i.Modules.DublinCore.Date != "" {
		pubDate = i.Modules.DublinCore.Date
	}

	var updated string
	if i.Modules.DublinCore.Modified != "" {
		updated = i.Modules.DublinCore.Modified
	}

	type thumbnail struct {
		URL    string `json:"url"`
		Width  int64  `json:"width"`
		Height int64  `json:"height"`
	}

	gt := thumbnail{
		URL:    i.Modules.Media.Thumbnail.URL,
		Width:  i.Modules.Media.Thumbnail.Width,
		Height: i.Modules.Media.Thumbnail.Height,
	}

	gi := &struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Link        string     `json:"link"`
		Description string     `json:"description"`
		Content     string     `json:"content"`
		PubDate     string     `json:"pubdate"`
		Updated     string     `json:"updated"`
		Authors     []author   `json:"authors"`
		Categories  []Category `json:"categories"`
		Enclosure   Enclosure  `json:"enclosure"`
		Thumbnail   thumbnail  `json:"thumbnail"`
	}{
		ID:          i.GUID.Value,
		Title:       i.Title,
		Link:        i.Link,
		Description: i.Description,
		Content:     i.Modules.Content.Encoded,
		PubDate:     pubDate,
		Updated:     updated,
		Authors:     authors,
		Categories:  i.Categories,
		Enclosure:   i.Enclosure,
		Thumbnail:   gt,
	}
	return json.Marshal(gi)
}
