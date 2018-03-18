package gss

import (
	"fmt"

	"github.com/naoto0822/gss/atom"
	"github.com/naoto0822/gss/rss1"
	"github.com/naoto0822/gss/rss2"
	"github.com/naoto0822/gss/session"
)

func NewClient() *Client {
	detector := newDetector()
	session := session.NewClient(nil)

	return &Client{
		detector: detector,
		session:  session,
	}
}

type Client struct {
	detector *detector
	session  *session.Client
}

func (c *Client) Feed(url string) (*Feed, error) {
	header := make(map[string]string)
	bytes, err := c.session.Get(url, header)
	if err != nil {
		return nil, err
	}

	rssType, err := c.detector.detect(bytes)
	if err != nil {
		return nil, err
	}

	parser, err := c.getParser(rssType)
	if err != nil {
		return nil, err
	}

	mappableFeed, err := parser.Parse(bytes)
	if err != nil {
		return nil, err
	}

	bytes, err := mappableFeed.ToJSON()
	if err != nil {
		return nil, err
	}

	feed := Feed{}
	err := feed.Map(bytes)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func (c *Client) getParser(rssType RSSType) (*interfaces.Parseable, error) {
	switch c {
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
