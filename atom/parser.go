package atom

import (
	"encoding/xml"
)

// Parser atom parser
type Parser struct {
}

// Parse run atom feed parsing
func (p *Parser) Parse(data []byte) (*Feed, error) {
	var feed Feed

	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}
