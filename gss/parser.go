package gss

import (
	"fmt"

	"github.com/naoto0822/gss/atom"
	"github.com/naoto0822/gss/interfaces"
	"github.com/naoto0822/gss/rss1"
	"github.com/naoto0822/gss/rss2"
)

func getParser(rssType RSSType) (interfaces.Parseable, error) {
	switch rssType {
	case RSS1:
		return rss1.NewParser(), nil
	case RSS2:
		return rss2.NewParser(), nil
	case Atom:
		return atom.NewParser(), nil
	default:
		return nil, fmt.Errorf("not found RSSType")
	}
}
