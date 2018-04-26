package modules

import ()

// DublinCore module
// support part of elemenet...
// see http://web.resource.org/rss/1.0/modules/dc/
type DublinCore struct {
	Title       string
	Creator     string
	Subject     string
	Description string
	Publisher   string
	Contributor string
	Date        string
	Rights      string
	Language    string
}
