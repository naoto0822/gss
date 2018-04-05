package rss2

import (
	"encoding/xml"

	"github.com/naoto0822/gss/interfaces"
)

// Parser RSS2.0 Parser
type Parser struct{}

// NewParser factory Parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse Rss2.0 feed parse
func (p *Parser) Parse(data []byte) (interfaces.Mappable, error) {
	var feed Feed

	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
