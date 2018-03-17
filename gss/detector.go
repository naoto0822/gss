package gss

import (
	"encoding/xml"
	"strings"
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

func (d *detector) detect(bytes []byte) (RSSType, error) {
	var root RootFeed
	xmlStr := string(bytes)

	// To Unmarshal root elements, gss attach dummy elements to target XML.
	processXMLStr := dummyStartElement + xmlStr + dummyEndElement

	err := xml.Unmarshal([]byte(processXMLStr), &root)
	if err != nil {
		return nil, err
	}

	feedType := Unknown

	if detector.isRSS1(root) {
		feedType = RSS1
	}

	if detector.isRSS2(root) {
		feedType = RSS2
	}

	if detector.isAtom(root) {
		feedType = Atom
	}

	return feedType, nil
}

func (d *detector) isRSS1(root RootFeed) bool {
	if root.RdfTag != "" && strings.Contains(root.RdfTag, "rss/1.0") {
		return true
	}
	return false
}

func (d *detector) isRSS2(root RootFeed) bool {
	if root.RssTag != "" && strings.Contains(root.RSSTag, "2.0") {
		return true
	}
	return false
}

func (d *detector) isAtom(root RootFeed) bool {
	if root.AtomTag != "" && strings.Contains(root.AtomTag, "Atom") {
		return true
	}
	return false
}
