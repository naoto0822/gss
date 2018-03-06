package gss

import (
	"encoding/xml"
	// "fmt"
)

const (
	dummyStartElement = "<root>"
	dummyEndElement   = "</root>"
)

type detector struct{}

func NewDetector() *detector {
	return &detector{}
}

func (d *detector) detect(bytes []byte) (*RootFeed, error) {
	var root RootFeed

	xmlStr := string(bytes)

	// To Unmarshal XML root elements, i attach dummy elements to target XML.
	processXMLStr := dummyStartElement + xmlStr + dummyEndElement

	err := xml.Unmarshal([]byte(processXMLStr), &root)
	if err != nil {
		return nil, err
	}

	return &root, nil
}
