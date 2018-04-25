package rss1

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/naoto0822/gss/modules"
)

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("TestNewParser not expected nil")
	}
}

func TestParseRSS1(t *testing.T) {
	path := "../testdata/rss_1.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestParseRSS1 ioutil.ReadFile returned error:", err)
	}

	parser := NewParser()
	feed, err := parser.Parse(bytes)

	if err != nil {
		t.Error("TestParseRSS1 Parse returned error:", err)
	}

	if feed == nil {
		t.Error("TestParseRSS1 Parse not expected nil")
	}

	// following expected Struct
	cDublinCore := modules.DublinCore{
		Date:     "2003-12-13T18:30:02Z",
		Language: "ja",
	}
	cModules := modules.Modules{
		DublinCore: cDublinCore,
	}

	channel := Channel{
		Title:       "Channel Title",
		Link:        "http://xml.com/pub",
		Description: "this is description.",
		Modules:     cModules,
	}

	image := Image{
		Title: "XML.com",
		Link:  "http://www.xml.com",
		URL:   "http://xml.com/universal/images/xml_tiny.gif",
	}

	item1DublinCore := modules.DublinCore{
		Date:    "2003-12-13T18:30:02Z",
		Creator: "記事1の作者名",
	}
	item1Modules := modules.Modules{
		DublinCore: item1DublinCore,
	}
	item1 := Item{
		Title:       "Processing Inclusions with XSLT",
		Link:        "http://xml.com/pub/2000/08/09/xslt/xslt.html",
		Description: "Processing document inclusions with general XML tools can be problematic. This article proposes a way of preserving inclusion information through SAX-based processing.",
		Modules:     item1Modules,
	}

	item2DublinCore := modules.DublinCore{
		Date:    "2003-12-13T18:30:02Z",
		Creator: "記事2の作者名",
	}
	item2Modules := modules.Modules{
		DublinCore: item2DublinCore,
	}
	item2 := Item{
		Title:       "Putting RDF to Work",
		Link:        "http://xml.com/pub/2000/08/09/rdfdb/index.html",
		Description: "Tool and API support for the Resource Description Framework is slowly coming of age. Edd Dumbill takes a look at RDFDB, one of the most exciting new RDF toolkits.",
		Modules:     item2Modules,
	}

	var items []Item
	items = append(items, item1)
	items = append(items, item2)

	textInput := TextInput{
		Title:       "Search XML.com",
		Description: "Search XML.com's XML collection",
		Name:        "s",
		Link:        "http://search.xml.com",
	}

	want := Feed{
		Channel:   channel,
		Image:     image,
		Items:     items,
		TextInput: textInput,
	}

	rss1Feed, ok := feed.(Feed)
	if !ok {
		t.Error("TestParseRSS1 not expected struct type")
	}

	if !reflect.DeepEqual(rss1Feed, want) {
		t.Error("TestParseRSS1 Parse not match Feed struct, ", feed, want)
	}
}

func TestParseErrorRSS1(t *testing.T) {
	path := "../testdata/rss_1.0_error.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestParseErrorRSS1 ioutil.ReadFile returned error:", err)
	}

	parser := NewParser()
	feed, err := parser.Parse(bytes)
	if err == nil {
		t.Error("TestParseErrorRSS1 Parse not expected return, feed:", feed)
	}
}
