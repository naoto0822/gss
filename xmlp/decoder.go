package xmlp

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// xmlp.Decoder (XML Pull Decoder)
// modify for reference mmcdole/goxpp.
// return xml.StartElement state.

// XMLTokenType is xml.Token Mapping
type XMLTokenType int

const (
	// StartXML is init state
	StartXML XMLTokenType = iota
	// EndXML is io.EOF state
	EndXML
	// StartElement is xml.StartElement
	StartElement
	// EndElement is xml.EndElement
	EndElement
	// CharData is xml.CharData
	CharData
	// Comment is xml.Comment
	Comment
	// ProcInst is xml.ProcInst
	ProcInst
	// Directive is xml.Directive
	Directive
)

// Decoder xml.Decoder client
type Decoder struct {
	TokenType XMLTokenType
	Space     string
	Local     string
	Attrs     []xml.Attr
	Text      string

	client *xml.Decoder
	token  xml.Token
}

// NewDecoder factory xmlp.Decoder
func NewDecoder(r io.Reader) *Decoder {
	client := xml.NewDecoder(r)
	return &Decoder{
		TokenType: StartXML,
		client:    client,
	}
}

// ElementName get namespace and local
func (d *Decoder) ElementName() string {
	if d.Space != "" {
		return d.Space + ":" + d.Local
	}
	return d.Local
}

// RootElement stating decode xml tree
func (d *Decoder) RootElement() error {
	for {
		err := d.next()
		if err != nil {
			return err
		}

		if d.TokenType == StartElement {
			break
		}

		if d.TokenType == EndXML {
			return fmt.Errorf("xmlp.StartRoot: failed finding root element before reaching XML last element")
		}
	}
	return nil
}

// NextElement iterate find next element
func (d *Decoder) NextElement() error {
	for {
		err := d.next()
		if err != nil {
			return err
		}

		if d.TokenType == StartElement || d.TokenType == EndElement {
			break
		}

		if d.TokenType == EndXML {
			return fmt.Errorf("xmlp.NextElement: failed finding next element before reaching XML last element")
		}
	}
	return nil
}

// DecodeElement decoding target element
func (d *Decoder) DecodeElement(v interface{}) error {
	if d.TokenType != StartElement {
		return fmt.Errorf("xmlp.DecodeElement: not StartElement")
	}

	token := d.token.(xml.StartElement)
	err := d.client.DecodeElement(v, &token)
	if err != nil {
		return err
	}

	local := d.Local
	d.resetState()
	d.TokenType = EndElement
	d.Local = local
	d.token = nil
	return nil
}

// Skip skip element
func (d *Decoder) Skip() error {
	for {
		err := d.currentToken()
		if err != nil {
			return err
		}

		if d.TokenType == StartElement {
			if err := d.Skip(); err != nil {
				return err
			}
		} else if d.TokenType == EndElement {
			return nil
		}
	}
}

// Expect to valid element in current token
func (d *Decoder) Expect(tokenType XMLTokenType, local string) error {
	if d.TokenType == tokenType && d.Local == local {
		return nil
	}
	return fmt.Errorf("xmlp.Expect: not expect current token")
}

func (d *Decoder) next() error {
	for {
		err := d.currentToken()
		if err != nil {
			return err
		}

		if d.TokenType == StartElement || d.TokenType == EndElement || d.TokenType == EndXML || d.TokenType == CharData {
			return nil
		}

		if d.TokenType == Comment || d.TokenType == ProcInst || d.TokenType == Directive {
			continue
		}
	}
}

func (d *Decoder) currentToken() error {
	d.resetState()

	token, err := d.client.Token()
	if err != nil {
		if err == io.EOF {
			d.token = nil
			d.TokenType = EndXML
			return nil
		}
		return err
	}

	d.token = xml.CopyToken(token)
	d.updateState(token)
	return nil
}

func (d *Decoder) updateState(t xml.Token) {
	d.setTokenType(t)

	switch tt := t.(type) {
	case xml.StartElement:
		d.updateStartElement(tt)
	case xml.EndElement:
		d.updateEndElement(tt)
	case xml.CharData:
		d.updateCharData(tt)
	}
}

func (d *Decoder) setTokenType(t xml.Token) {
	switch t.(type) {
	case xml.StartElement:
		d.TokenType = StartElement
	case xml.EndElement:
		d.TokenType = EndElement
	case xml.CharData:
		d.TokenType = CharData
	case xml.Comment:
		d.TokenType = Comment
	case xml.ProcInst:
		d.TokenType = ProcInst
	case xml.Directive:
		d.TokenType = Directive
	}
}

func (d *Decoder) updateStartElement(t xml.StartElement) {
	d.Space = strings.ToLower(t.Name.Space)
	d.Local = strings.ToLower(t.Name.Local)
	d.Attrs = t.Attr
}

func (d *Decoder) updateEndElement(t xml.EndElement) {
	d.Local = strings.ToLower(t.Name.Local)
}

func (d *Decoder) updateCharData(t xml.CharData) {
	d.Text = string([]byte(t))
}

func (d *Decoder) resetState() {
	d.Space = ""
	d.Local = ""
	d.Attrs = []xml.Attr{}
	d.Text = ""
}
