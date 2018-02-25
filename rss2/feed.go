package rss2

import ()

// Feed RSS2.0 feed
type Feed struct {
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
	Category       []Category `xml:"category"`
	Generator      string     `xml:"generator"`
	Docs           string     `xml:"docs"`
	Cloud          Cloud      `xml:"cloud"`
	TTL            int        `xml:"ttl"`
	Image          Image      `xml:"image"`
	Rating         string     `xml:"rating"`
	// TODO: continue
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
