package gss

import ()

// rootFeed is detected struct
type RootFeed struct {
	XMLTag  XMLTag  `xml:"xml"`
	RdfTag  rdfTag  `xml:"RDF"`
	RSSTag  rssTag  `xml:"rss"`
	AtomTag atomTag `xml:"feed"`
}

// XMLTag
type XMLTag struct {
	Version string `xml:"version,attr"`
}

// rdfTag is RSS1.0 root elements
type rdfTag struct {
	Xmlns string `xml:"xmlns,attr"`
}

// rssTag is RSS2.0 root elements
type rssTag struct {
	Version string `xml:"version,attr"`
}

// atomTag is Atom root elements
type atomTag struct {
	Xmlns string `xml:"xmlns,attr"`
}
