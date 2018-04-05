package atom

import (
	"encoding/xml"

	"github.com/naoto0822/gss/interfaces"
)

// Parser atom parser
type Parser struct {
	interfaces.Parseable
}

// NewParser factory Parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse run atom feed parsing
func (p *Parser) Parse(data []byte) (interfaces.Mappable, error) {
	var feed Feed

	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
