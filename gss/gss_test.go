package gss

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Error("TestNewClient fail factory client")
	}
}

func TestFeedRSS1(t *testing.T) {
	normalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/rss_1.0.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(normalHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedRSS1 fail get RSS response")
	}

	// assemble expect feed
	a1 := Author{
		Name: "記事1の作者名",
	}
	authors1 := []Author{a1}
	item1 := Item{
		Title:       "Processing Inclusions with XSLT",
		Link:        "http://xml.com/pub/2000/08/09/xslt/xslt.html",
		Description: "Processing document inclusions with general XML tools can be problematic. This article proposes a way of preserving inclusion information through SAX-based processing.",
		PubDate:     "2003-12-13T18:30:02Z",
		Authors:     authors1,
	}

	a2 := Author{
		Name: "記事2の作者名",
	}
	authors2 := []Author{a2}
	item2 := Item{
		Title:       "Putting RDF to Work",
		Link:        "http://xml.com/pub/2000/08/09/rdfdb/index.html",
		Description: "Tool and API support for the Resource Description Framework is slowly coming of age. Edd Dumbill takes a look at RDFDB, one of the most exciting new RDF toolkits.",
		PubDate:     "2003-12-13T18:30:02Z",
		Authors:     authors2,
	}

	items := []Item{item1, item2}

	image := Image{
		Title: "XML.com",
		Link:  "http://www.xml.com",
		URL:   "http://xml.com/universal/images/xml_tiny.gif",
	}

	want := &Feed{
		Title:       "Channel Title",
		Link:        "http://xml.com/pub",
		Description: "this is description.",
		Image:       image,
		PubDate:     "2003-12-13T18:30:02Z",
		Items:       items,
	}

	if !reflect.DeepEqual(ret.Feed, want) {
		t.Error("TestFeedRSS1 not match expect feed, ", ret.Feed, want)
	}

	if ret.RSSType != RSS1 {
		t.Error("TestFeedRSS1 not expected rss type, ", ret.RSSType)
	}

	if ret.RSS1Feed == nil || ret.RSS2Feed != nil || ret.AtomFeed != nil {
		t.Error("TestFeedRSS1 failed assemble Result")
	}
}

func TestErrorFeedRSS1(t *testing.T) {
	errorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/rss_1.0_error.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(errorHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)
	if ret != nil || err == nil {
		t.Error("TestErrorFeedRSS1 not expected return")
	}
}

func TestFeedRSS2(t *testing.T) {
	normalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/rss_2.0.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(normalHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedRSS2 fail get RSS response")
	}

	// assemble expecting feed
	item1 := Item{
		ID:          "http://liftoff.msfc.nasa.gov/2003/06/03.html#item573",
		Title:       "Star City",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp",
		Description: `How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.`,
		PubDate:     "Tue, 03 Jun 2003 09:39:21 GMT",
		Content:     "This is <i>italics</i>.",
	}

	thumbnail2 := Thumbnail{
		URL:    "http://www.foo.com/keyframe.jpg",
		Width:  75,
		Height: 50,
	}
	item2 := Item{
		ID:          "http://liftoff.msfc.nasa.gov/2003/05/30.html#item572",
		Description: `this is <b>bold</b>`,
		PubDate:     "Fri, 30 May 2003 11:06:42 GMT",
		Thumbnail:   thumbnail2,
	}

	item3 := Item{
		ID:          "http://liftoff.msfc.nasa.gov/2003/05/27.html#item571",
		Title:       "The Engine That Does More",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp",
		Description: `Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly. The proposed VASIMR engine would do that.`,
		PubDate:     "Tue, 27 May 2003 08:37:32 GMT",
		Content:     "This is <b>bold</b>.",
	}

	item4 := Item{
		ID:          "http://liftoff.msfc.nasa.gov/2003/05/20.html#item570",
		Title:       "Astronauts' Dirty Laundry",
		Link:        "http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp",
		Description: `Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them. Instead, astronauts have other options.`,
		PubDate:     "Tue, 20 May 2003 08:56:02 GMT",
	}

	items := []Item{item1, item2, item3, item4}

	want := &Feed{
		Title:       "Liftoff News",
		Link:        "http://liftoff.msfc.nasa.gov/",
		Description: "Liftoff to Space Exploration.",
		PubDate:     "Tue, 10 Jun 2003 04:00:00 GMT",
		Updated:     "Tue, 10 Jun 2003 09:41:01 GMT",
		Items:       items,
	}

	if !reflect.DeepEqual(ret.Feed, want) {
		t.Error("TestFeedRSS2 not match expect feed, ", ret.Feed, want)
	}

	if ret.RSSType != RSS2 {
		t.Error("TestFeedRSS2 not expected rss type, ", ret.RSSType)
	}

	if ret.RSS1Feed != nil || ret.RSS2Feed == nil || ret.AtomFeed != nil {
		t.Error("TestFeedRSS2 failed assemble Result")
	}
}

func TestErrorFeedRss2(t *testing.T) {
	errorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/rss_2.0_error.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(errorHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)
	if ret != nil || err == nil {
		t.Error("TestErrorFeedRss2 not expected return")
	}
}

func TestFeedAtom(t *testing.T) {
	normalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/atom_1.0.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(normalHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedAtom fail get RSS response")
	}

	// assemble expecting feed
	item1 := Item{
		ID:          "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:       "This is Atom 1.0",
		Link:        "http://example.org/2003/12/13/atom03",
		Description: "Some text.",
		Updated:     "2003-12-13T18:30:02Z",
		PubDate:     "2003-12-13T08:29:29-04:00",
	}

	categories2 := []string{"Music", "Sports"}
	item2 := Item{
		ID:         "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:      "Today is hot",
		Link:       "http://example.org/2003/12/13/atom03",
		Content:    "This is <b>red&blue</b>.",
		Updated:    "2003-12-13T18:30:02Z",
		PubDate:    "2018-12-13T08:29:29-04:00",
		Categories: categories2,
	}

	items := []Item{item1, item2}

	a := Author{
		Name:  "John Doe",
		Email: "JohnDoe@example.com",
	}
	authors := []Author{a}

	image := Image{
		URL: "logo.jpg",
	}

	want := &Feed{
		Title:       "Example Feed",
		Link:        "http://google.com",
		Description: "This is Sub title",
		Image:       image,
		Updated:     "2003-12-13T18:30:02Z",
		Authors:     authors,
		CopyRight:   "2005 naoto0822",
		Items:       items,
	}

	if !reflect.DeepEqual(ret.Feed, want) {
		t.Error("TestFeedAtom not match expect feed, ", ret.Feed, want)
	}

	if ret.RSSType != Atom {
		t.Error("TestFeedAtom not expected rss type, ", ret.RSSType)
	}

	if ret.RSS1Feed != nil || ret.RSS2Feed != nil || ret.AtomFeed == nil {
		t.Error("TestFeedAtom failed assemble Result")
	}
}

func TestErrorFeedAtom(t *testing.T) {
	errorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../testdata/atom_1.0_error.xml"
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})

	testServer := httptest.NewServer(errorHandler)
	defer testServer.Close()

	client := NewClient()
	ret, err := client.Parse(testServer.URL)

	if ret != nil || err == nil {
		t.Error("TestErrorFeedAtom not expected return")
	}
}
