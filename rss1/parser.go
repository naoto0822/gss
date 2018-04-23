package rss1

import (
	"bytes"

	"github.com/naoto0822/gss/interfaces"
	"github.com/naoto0822/gss/xmlp"
)

// Parser RSS1.0 parser
type Parser struct{}

// NewParser factory Parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse RSS1.0 feed parse
func (p *Parser) Parse(data []byte) (interfaces.Mappable, error) {
	r := bytes.NewReader(data)
	decoder := xmlp.NewDecoder(r)
	err := decoder.RootElement()
	if err != nil {
		return nil, err
	}

	var feed Feed
	err = p.decode(decoder, &feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (p *Parser) decode(d *xmlp.Decoder, f *Feed) error {
	if err := d.Expect(xmlp.StartElement, "rdf"); err != nil {
		return err
	}

	for {
		err := d.NextElement()
		if err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			switch d.Local {
			case "channel":
				var channel Channel
				err := p.decodeChannel(d, &channel)
				if err != nil {
					return err
				}
				f.Channel = channel
			case "image":
				var image Image
				err := d.DecodeElement(&image)
				if err != nil {
					return err
				}
				f.Image = image
			case "item":
				var item Item
				err := p.decodeItem(d, &item)
				if err != nil {
					return err
				}
				f.Items = append(f.Items, item)
			case "textinput":
				var textInput TextInput
				err := d.DecodeElement(&textInput)
				if err != nil {
					return err
				}
				f.TextInput = textInput
			default:
				err := d.Skip()
				if err != nil {
					return err
				}
			}
		}
	}

	if err := d.Expect(xmlp.EndElement, "rdf"); err != nil {
		return err
	}

	return nil
}

func (p *Parser) decodeChannel(d *xmlp.Decoder, c *Channel) error {
	if err := d.Expect(xmlp.StartElement, "channel"); err != nil {
		return err
	}

	for {
		err := d.NextElement()
		if err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			switch d.Local {
			case "title":
				var title string
				err := d.DecodeElement(&title)
				if err != nil {
					return err
				}
				c.Title = title
			case "link":
				var link string
				err := d.DecodeElement(&link)
				if err != nil {
					return err
				}
				c.Link = link
			case "description":
				var description string
				err := d.DecodeElement(&description)
				if err != nil {
					return err
				}
				c.Description = description
			case "date":
				var date string
				err := d.DecodeElement(&date)
				if err != nil {
					return err
				}
				c.Date = date
			case "language":
				var lang string
				err := d.DecodeElement(&lang)
				if err != nil {
					return err
				}
				c.Language = lang
			default:
				err := d.Skip()
				if err != nil {
					return err
				}
			}
		}
	}

	if err := d.Expect(xmlp.EndElement, "channel"); err != nil {
		return err
	}

	return nil
}

func (p *Parser) decodeItem(d *xmlp.Decoder, i *Item) error {
	if err := d.Expect(xmlp.StartElement, "item"); err != nil {
		return err
	}

	for {
		err := d.NextElement()
		if err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			switch d.Local {
			case "title":
				var title string
				err := d.DecodeElement(&title)
				if err != nil {
					return err
				}
				i.Title = title
			case "link":
				var link string
				err := d.DecodeElement(&link)
				if err != nil {
					return err
				}
				i.Link = link
			case "description":
				var description string
				err := d.DecodeElement(&description)
				if err != nil {
					return err
				}
				i.Description = description
			case "creator":
				var creator string
				err := d.DecodeElement(&creator)
				if err != nil {
					return err
				}
				i.Creator = creator
			case "date":
				var date string
				err := d.DecodeElement(&date)
				if err != nil {
					return err
				}
				i.Date = date
			default:
				err := d.Skip()
				if err != nil {
					return err
				}
			}
		}
	}

	if err := d.Expect(xmlp.EndElement, "item"); err != nil {
		return err
	}

	return nil
}
