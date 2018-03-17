package gss

// RSSType is identifies Feed
type RSSType int

const (
	// Unknown can't detected feed type
	Unknown RSSType = iota
	// RSS1 RSS Version 1.0
	RSS1
	// RSS2 RSS Version 2.0
	RSS2
	// Atom Format is Atom
	Atom
)
