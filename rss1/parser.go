package rss1

import (
	"encoding/xml"
)

// Parser RSS1.0 parser
type Parser struct{}

// NewParser factory Parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse RSS1.0 feed parse
func (p *Parser) Parse(data []byte) (*Feed, error) {
	var feed Feed

	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}