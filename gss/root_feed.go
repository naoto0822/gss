package gss

// rootFeed is detected struct
type rootFeed struct {
	RdfTag  rdfTag  `xml:"RDF"`
	RssTag  rssTag  `xml:"rss"`
	AtomTag atomTag `xml:"feed"`
	RssType RssType
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
