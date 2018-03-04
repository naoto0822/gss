package gss

// RssType is identifies Feed
type RssType int

const (
	UnknownType RssType = iota
	Rss1Type
	Rss2Type
	AtomType
)
