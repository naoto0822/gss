package modules

// Media RSS Module
// support part of elemenet...
// see http://www.rssboard.org/media-rss
type Media struct {
	Thumbnail mediaThumbnail `xml:"thumbnail"`
}

type mediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int64  `xml:"width,attr"`
	Height int64  `xml:"height,attr"`
}
