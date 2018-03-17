package gss

import ()

// rootFeed is detected struct
type rootFeed struct {
	XMLTag  xmlTag  `xml:"xml"`
	RdfTag  rdfTag  `xml:"RDF"`
	RSSTag  rssTag  `xml:"rss"`
	AtomTag atomTag `xml:"feed"`
}

type xmlTag struct {
	Version string `xml:"version,attr"`
}

type rdfTag struct {
	Xmlns string `xml:"xmlns,attr"`
}

type rssTag struct {
	Version string `xml:"version,attr"`
}

type atomTag struct {
	Xmlns string `xml:"xmlns,attr"`
}
