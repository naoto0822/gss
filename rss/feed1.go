package rss

import ()

// Feed1 RSS1.0 feed
type Feed1 struct {
	Channel Channel `xml:"channel"`
	Image   Image   `xml:"image"`
	Items   []Item  `xml:"item"`
}

// Channel RSS1.0 channel
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Date        string `xml:"dc:date"`
	Language    string `xml:"dc:language"`
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
}
