package gss

import (
	"github.com/naoto0822/gss/atom"
	"github.com/naoto0822/gss/interfaces"
	"github.com/naoto0822/gss/rss1"
	"github.com/naoto0822/gss/rss2"
)

// Result is returned by gss.Client.Parse
type Result struct {
	RSSType  RSSType
	Feed     *Feed
	RSS1Feed *rss1.Feed
	RSS2Feed *rss2.Feed
	AtomFeed *atom.Feed
}

func makeResult(rssType RSSType, feed Feed, rawFeed interfaces.Mappable) Result {
	result := Result{
		RSSType: rssType,
		Feed:    &feed,
	}

	switch raw := rawFeed.(type) {
	case rss1.Feed:
		result.RSS1Feed = &raw
	case rss2.Feed:
		result.RSS2Feed = &raw
	case atom.Feed:
		result.AtomFeed = &raw
	}

	return result
}
