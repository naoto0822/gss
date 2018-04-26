package modules

// Media RSS Module
// support part of elemenet...
// see http://www.rssboard.org/media-rss
type Media struct {
	Thumbnail MediaThumbnail `xml:"thumbnail"`
}

// MediaThumbnail Media Thumbnail Module
type MediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int64  `xml:"width,attr"`
	Height int64  `xml:"height,attr"`
}
