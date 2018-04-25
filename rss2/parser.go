package rss2

import (
	"bytes"
	//"fmt"

	"github.com/naoto0822/gss/interfaces"
	"github.com/naoto0822/gss/modules"
	"github.com/naoto0822/gss/xmlp"
)

// Parser RSS2.0 Parser
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

// Parse RSS2.0 feed parse
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

	if err := d.Expect(xmlp.StartElement, "rss"); err != nil {
		return err
	}

	for {
		if err := d.NextElement(); err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			local := d.Local
			isModule := p.moduleDecoder.IsModule(d)

			if local == "channel" && !isModule {
				var channel Channel
				if err := p.decodeChannel(d, &channel); err != nil {
					return err
				}
				f.Channel = channel
			} else {
				if err := d.Skip(); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func (p *Parser) decodeChannel(d *xmlp.Decoder, c *Channel) error {
	if err := d.Expect(xmlp.StartElement, "channel"); err != nil {
		return err
	}

	for {
		if err := d.NextElement(); err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			local := d.Local
			isModule := p.moduleDecoder.IsModule(d)

			if local == "title" && !isModule {
				var title string
				if err := d.DecodeElement(&title); err != nil {
					return err
				}
				c.Title = title
			} else if local == "link" && !isModule {
				var link string
				if err := d.DecodeElement(&link); err != nil {
					return err
				}
				c.Link = link
			} else if local == "description" && !isModule {
				var description string
				if err := d.DecodeElement(&description); err != nil {
					return err
				}
				c.Description = description
			} else if local == "language" && !isModule {
				var lang string
				if err := d.DecodeElement(&lang); err != nil {
					return err
				}
				c.Language = lang
			} else if local == "copyright" && !isModule {
				var copyright string
				if err := d.DecodeElement(&copyright); err != nil {
					return err
				}
				c.CopyRight = copyright
			} else if local == "managingeditor" && !isModule {
				var editor string
				if err := d.DecodeElement(&editor); err != nil {
					return err
				}
				c.ManagingEditor = editor
			} else if local == "webmaster" && !isModule {
				var webMaster string
				if err := d.DecodeElement(&webMaster); err != nil {
					return err
				}
				c.WebMaster = webMaster
			} else if local == "pubdate" && !isModule {
				var pubDate string
				if err := d.DecodeElement(&pubDate); err != nil {
					return err
				}
				c.PubDate = pubDate
			} else if local == "lastbuilddate" && !isModule {
				var date string
				if err := d.DecodeElement(&date); err != nil {
					return err
				}
				c.LastBuildDate = date
			} else if local == "category" && !isModule {
				var category Category
				if err := d.DecodeElement(&category); err != nil {
					return err
				}
				c.Categories = append(c.Categories, category)
			} else if local == "generator" && !isModule {
				var generator string
				if err := d.DecodeElement(&generator); err != nil {
					return err
				}
				c.Generator = generator
			} else if local == "docs" && !isModule {
				var docs string
				if err := d.DecodeElement(&docs); err != nil {
					return err
				}
				c.Docs = docs
			} else if local == "cloud" && !isModule {
				var cloud Cloud
				if err := d.DecodeElement(&cloud); err != nil {
					return err
				}
				c.Cloud = cloud
			} else if local == "ttl" && !isModule {
				var ttl int
				if err := d.DecodeElement(&ttl); err != nil {
					return err
				}
				c.TTL = ttl
			} else if local == "image" && !isModule {
				var image Image
				if err := d.DecodeElement(&image); err != nil {
					return err
				}
				c.Image = image
			} else if local == "rating" && !isModule {
				var rating string
				if err := d.DecodeElement(&rating); err != nil {
					return err
				}
				c.Rating = rating
			} else if local == "textinput" && !isModule {
				var textInput TextInput
				if err := d.DecodeElement(&textInput); err != nil {
					return err
				}
				c.TextInput = textInput
			} else if local == "skiphours" && !isModule {
				var hours SkipHours
				if err := d.DecodeElement(&hours); err != nil {
					return err
				}
				c.SkipHours = hours
			} else if local == "skipdays" && !isModule {
				var days SkipDays
				if err := d.DecodeElement(&days); err != nil {
					return err
				}
				c.SkipDays = days
			} else if local == "item" && !isModule {
				var item Item
				if err := p.decodeItem(d, &item); err != nil {
					return err
				}
				c.Items = append(c.Items, item)
			} else if isModule {
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
		if err := d.NextElement(); err != nil {
			return err
		}

		if d.TokenType == xmlp.EndElement {
			break
		}

		switch d.TokenType {
		case xmlp.StartElement:
			local := d.Local
			isModule := p.moduleDecoder.IsModule(d)

			if local == "title" && !isModule {
				var title string
				if err := d.DecodeElement(&title); err != nil {
					return err
				}
				i.Title = title
			} else if local == "link" && !isModule {
				var link string
				if err := d.DecodeElement(&link); err != nil {
					return err
				}
				i.Link = link
			} else if local == "description" && !isModule {
				var description string
				if err := d.DecodeElement(&description); err != nil {
					return err
				}
				i.Description = description
			} else if local == "author" && !isModule {
				var author string
				if err := d.DecodeElement(&author); err != nil {
					return err
				}
				i.Author = author
			} else if local == "category" && !isModule {
				var category Category
				if err := d.DecodeElement(&category); err != nil {
					return err
				}
				i.Categories = append(i.Categories, category)
			} else if local == "comments" && !isModule {
				var comments string
				if err := d.DecodeElement(&comments); err != nil {
					return err
				}
				i.Comments = comments
			} else if local == "enclosure" && !isModule {
				var enclosure Enclosure
				if err := d.DecodeElement(&enclosure); err != nil {
					return err
				}
				i.Enclosure = enclosure
			} else if local == "guid" && !isModule {
				var guid GUID
				if err := d.DecodeElement(&guid); err != nil {
					return err
				}
				i.GUID = guid
			} else if local == "pubdate" && !isModule {
				var pubDate string
				if err := d.DecodeElement(&pubDate); err != nil {
					return err
				}
				i.PubDate = pubDate
			} else if local == "source" && !isModule {
				var source Source
				if err := d.DecodeElement(&source); err != nil {
					return err
				}
				i.Source = source
			} else if isModule {
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
