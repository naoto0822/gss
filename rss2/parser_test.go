package rss2

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

func TestParseRSS2(t *testing.T) {
	path := "../testdata/rss_2.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestParseRSS2 ioutil.ReadFile returned error:", err)
	}

	parser := NewParser()
	feed, err := parser.Parse(bytes)

	if err != nil {
		t.Error("TestParseRSS2 Parse returned error:", err)
	}

	if feed == nil {
		t.Error("TestParseRSS2 Parse not expected nil")
	}

	// following expected Struct
	guid1 := GUID{
		Value: "http://liftoff.msfc.nasa.gov/2003/06/03.html#item573",
	}

	desc1 := `How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.`

	contentModule1 := modules.Content{
		Encoded: "This is <i>italics</i>.",
	}
	module1 := modules.Modules{
		Content: contentModule1,
	}

	item1 := Item{
		Title:       "Star City",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp",
		Description: desc1,
		PubDate:     "Tue, 03 Jun 2003 09:39:21 GMT",
		GUID:        guid1,
		Modules:     module1,
	}

	guid2 := GUID{
		Value: "http://liftoff.msfc.nasa.gov/2003/05/30.html#item572",
	}

	desc2 := `this is <b>bold</b>`

	thumbnail2 := modules.MediaThumbnail{
		URL:    "http://www.foo.com/keyframe.jpg",
		Width:  75,
		Height: 50,
	}
	mediaModule2 := modules.Media{
		Thumbnail: thumbnail2,
	}
	modules2 := modules.Modules{
		Media: mediaModule2,
	}

	item2 := Item{
		Description: desc2,
		PubDate:     "Fri, 30 May 2003 11:06:42 GMT",
		GUID:        guid2,
		Modules:     modules2,
	}

	guid3 := GUID{
		Value: "http://liftoff.msfc.nasa.gov/2003/05/27.html#item571",
	}

	contentModule3 := modules.Content{
		Encoded: "This is <b>bold</b>.",
	}
	modules3 := modules.Modules{
		Content: contentModule3,
	}

	item3 := Item{
		Title:       "The Engine That Does More",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp",
		Description: "Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly. The proposed VASIMR engine would do that.",
		PubDate:     "Tue, 27 May 2003 08:37:32 GMT",
		GUID:        guid3,
		Modules:     modules3,
	}

	guid4 := GUID{
		Value: "http://liftoff.msfc.nasa.gov/2003/05/20.html#item570",
	}

	item4 := Item{
		Title:       "Astronauts' Dirty Laundry",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp",
		Description: "Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them. Instead, astronauts have other options.",
		PubDate:     "Tue, 20 May 2003 08:56:02 GMT",
		GUID:        guid4,
	}

	var items []Item
	items = append(items, item1)
	items = append(items, item2)
	items = append(items, item3)
	items = append(items, item4)

	channel := Channel{
		Title:          "Liftoff News",
		Link:           "http://liftoff.msfc.nasa.gov/",
		Description:    "Liftoff to Space Exploration.",
		Language:       "en-us",
		PubDate:        "Tue, 10 Jun 2003 04:00:00 GMT",
		LastBuildDate:  "Tue, 10 Jun 2003 09:41:01 GMT",
		Docs:           "http://blogs.law.harvard.edu/tech/rss",
		Generator:      "Weblog Editor 2.0",
		ManagingEditor: "editor@example.com",
		WebMaster:      "webmaster@example.com",
		Items:          items,
	}

	want := Feed{
		Channel: channel,
	}

	rss2Feed, ok := feed.(Feed)
	if !ok {
		t.Error("TestParseRSS2 not expected type")
	}

	if !reflect.DeepEqual(rss2Feed, want) {
		t.Error("TestParseRSS2 not match, ", feed, want)
	}
}

func TestParseErrorRSS2(t *testing.T) {
	path := "../testdata/rss_2.0_error.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestParseErrorRSS2 ioutil.ReadFile returned error:", err)
	}

	parser := NewParser()
	feed, err := parser.Parse(bytes)
	if err == nil {
		t.Error("TestParseErrorRSS2 Parse not expected return, feed:", feed)
	}
}
