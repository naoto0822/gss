package rss2

import (
	"html/template"
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
	Copyright      string     `xml:"copyright"`
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
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	Author      string        `xml:"author"`
	Categories  []Category    `xml:"category"`
	Comments    string        `xml:"comments"`
	Enclosure   Enclosure     `xml:"enclosure"`
	GUID        GUID          `xml:"guid"`
	PubDate     string        `xml:"pubDate"`
	Source      Source        `xml:"source"`
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
