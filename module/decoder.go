package module

import (
	"github.com/naoto0822/gss/xmlp"
)

const (
	dublinCoreSpace = "http://purl.org/dc/elements/1.1/"
	mediaSpace      = ""
	contentSpace    = ""
)

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (md *Decoder) IsModule(d *xmlp.Decoder) bool {
	space := d.Space
	if space == dublinCoreSpace {
		return true
	}
	return false
}

func (md *Decoder) DecodeElement(d *xmlp.Decoder, m *Modules) error {
	space := d.Space

	switch space {
	case dublinCoreSpace:
		if err := md.decodeDublinCore(d, &m.DublinCore); err != nil {
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
	default:
		if err := d.Skip(); err != nil {
			return err
		}
	}

	return nil
}
