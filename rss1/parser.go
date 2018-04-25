package rss1

import (
	"bytes"

	"github.com/naoto0822/gss/interfaces"
	"github.com/naoto0822/gss/modules"
	"github.com/naoto0822/gss/xmlp"
)

// Parser RSS1.0 parser
type Parser struct {
	moduleDecoder *modules.Decoder
}

// NewParser factory Parser
func NewParser() *Parser {
	md := modules.NewDecoder()
	return &Parser{
		moduleDecoder: md,
	}
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
			local := d.Local

			if local == "channel" {
				var channel Channel
				if err := p.decodeChannel(d, &channel); err != nil {
					return err
				}
				f.Channel = channel
			} else if local == "image" {
				var image Image
				if err := d.DecodeElement(&image); err != nil {
					return err
				}
				f.Image = image
			} else if local == "item" {
				var item Item
				if err := p.decodeItem(d, &item); err != nil {
					return err
				}
				f.Items = append(f.Items, item)
			} else if local == "textinput" {
				var textInput TextInput
				if err := d.DecodeElement(&textInput); err != nil {
					return err
				}
				f.TextInput = textInput
			} else {
				if err := d.Skip(); err != nil {
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
			local := d.Local

			if local == "title" {
				var title string
				if err := d.DecodeElement(&title); err != nil {
					return err
				}
				c.Title = title
			} else if local == "link" {
				var link string
				if err := d.DecodeElement(&link); err != nil {
					return err
				}
				c.Link = link
			} else if local == "description" {
				var description string
				if err := d.DecodeElement(&description); err != nil {
					return err
				}
				c.Description = description
			} else if p.moduleDecoder.IsModule(d) {
				if err := p.moduleDecoder.DecodeElement(d, &c.Modules); err != nil {
					return err
				}
			} else {
				if err := d.Skip(); err != nil {
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
			local := d.Local

			if local == "title" {
				var title string
				if err := d.DecodeElement(&title); err != nil {
					return err
				}
				i.Title = title
			} else if local == "link" {
				var link string
				if err := d.DecodeElement(&link); err != nil {
					return err
				}
				i.Link = link
			} else if local == "description" {
				var description string
				if err := d.DecodeElement(&description); err != nil {
					return err
				}
				i.Description = description
			} else if p.moduleDecoder.IsModule(d) {
				if err := p.moduleDecoder.DecodeElement(d, &i.Modules); err != nil {
					return err
				}
			} else {
				if err := d.Skip(); err != nil {
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
