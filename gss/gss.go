package gss

import (
	"github.com/naoto0822/gss/session"
)

// NewClient factory Client
func NewClient() *Client {
	detector := newDetector()
	session := session.NewClient(nil)

	return &Client{
		detector: detector,
		session:  session,
	}
}

// Client gss Client
type Client struct {
	detector *detector
	session  *session.Client
}

// Parse get gss.Result
func (c *Client) Parse(url string) (*Result, error) {
	header := make(map[string]string)
	bytes, err := c.session.Get(url, header)
	if err != nil {
		return nil, err
	}

	rssType, err := c.detector.detect(bytes)
	if err != nil {
		return nil, err
	}

	parser, err := getParser(rssType)
	if err != nil {
		return nil, err
	}

	mappableFeed, err := parser.Parse(bytes)
	if err != nil {
		return nil, err
	}

	feedBytes, err := mappableFeed.ToJSON()
	if err != nil {
		return nil, err
	}

	feed := Feed{}
	err = feed.Map(feedBytes)
	if err != nil {
		return nil, err
	}

	ret := makeResult(rssType, feed, mappableFeed)
	return &ret, nil
}
