package modules

import (
	"github.com/naoto0822/gss/xmlp"
)

const (
	dublinCoreSpace = "http://purl.org/dc/elements/1.1/"
	mediaSpace      = "http://search.yahoo.com/mrss/"
	contentSpace    = "http://purl.org/rss/1.0/modules/content/"
)

// Decoder RSS modules decoder
type Decoder struct{}

// NewDecoder factory module Decoder
func NewDecoder() *Decoder {
	return &Decoder{}
}

// IsModule current Element is RSS module?
func (md *Decoder) IsModule(d *xmlp.Decoder) bool {
	space := d.Space
	if space == dublinCoreSpace || space == mediaSpace || space == contentSpace {
		return true
	}
	return false
}

// DecodeElement RSS module decode
func (md *Decoder) DecodeElement(d *xmlp.Decoder, m *Modules) error {
	space := d.Space

	switch space {
	case dublinCoreSpace:
		if err := md.decodeDublinCore(d, &m.DublinCore); err != nil {
			return err
		}
	case mediaSpace:
		if err := md.decodeMedia(d, &m.Media); err != nil {
			return err
		}
	case contentSpace:
		if err := md.decodeContent(d, &m.Content); err != nil {
			return err
		}
	default:
		if err := d.Skip(); err != nil {
			return err
		}
	}

	return nil
}

func (md *Decoder) decodeDublinCore(d *xmlp.Decoder, dc *DublinCore) error {
	local := d.Local

	switch local {
	case "title":
		var title string
		if err := d.DecodeElement(&title); err != nil {
			return err
		}
		dc.Title = title
	case "creator":
		var creator string
		if err := d.DecodeElement(&creator); err != nil {
			return err
		}
		dc.Creator = creator
	case "subject":
		var subject string
		if err := d.DecodeElement(&subject); err != nil {
			return err
		}
		dc.Subject = subject
	case "description":
		var description string
		if err := d.DecodeElement(&description); err != nil {
			return err
		}
		dc.Description = description
	case "publisher":
		var publisher string
		if err := d.DecodeElement(&publisher); err != nil {
			return err
		}
		dc.Publisher = publisher
	case "date":
		var date string
		if err := d.DecodeElement(&date); err != nil {
			return err
		}
		dc.Date = date
	case "rights":
		var rights string
		if err := d.DecodeElement(&rights); err != nil {
			return err
		}
		dc.Rights = rights
	case "language":
		var lang string
		if err := d.DecodeElement(&lang); err != nil {
			return err
		}
		dc.Language = lang
	case "modified":
		var modified string
		if err := d.DecodeElement(&modified); err != nil {
			return err
		}
		dc.Modified = modified
	default:
		if err := d.Skip(); err != nil {
			return err
		}
	}

	return nil
}

func (md *Decoder) decodeMedia(d *xmlp.Decoder, m *Media) error {
	local := d.Local

	switch local {
	case "thumbnail":
		var thumbnail MediaThumbnail
		if err := d.DecodeElement(&thumbnail); err != nil {
			return err
		}
		m.Thumbnail = thumbnail
	default:
		if err := d.Skip(); err != nil {
			return err
		}
	}

	return nil
}

func (md *Decoder) decodeContent(d *xmlp.Decoder, c *Content) error {
	local := d.Local

	switch local {
	case "encoded":
		var encoded string
		if err := d.DecodeElement(&encoded); err != nil {
			return err
		}
		c.Encoded = encoded
	default:
		if err := d.Skip(); err != nil {
			return err
		}
	}

	return nil
}
