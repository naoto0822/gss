package gss

import (
	//"github.com/naoto0822/gss/session"
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
	feed, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedRSS1 fail get RSS response")
	}

	// assemble expect feed
	a1 := Author{
		Name: "記事1の作者名",
	}
	authors1 := []Author{a1}
	links1 := []string{"http://xml.com/pub/2000/08/09/xslt/xslt.html"}
	item1 := Item{
		Title:   "Processing Inclusions with XSLT",
		Links:   links1,
		Body:    "Processing document inclusions with general XML tools can be problematic. This article proposes a way of preserving inclusion information through SAX-based processing.",
		PubDate: "2003-12-13T18:30:02Z",
		Authors: authors1,
	}

	a2 := Author{
		Name: "記事2の作者名",
	}
	authors2 := []Author{a2}
	links2 := []string{"http://xml.com/pub/2000/08/09/rdfdb/index.html"}
	item2 := Item{
		Title:   "Putting RDF to Work",
		Links:   links2,
		Body:    "Tool and API support for the Resource Description Framework is slowly coming of age. Edd Dumbill takes a look at RDFDB, one of the most exciting new RDF toolkits.",
		PubDate: "2003-12-13T18:30:02Z",
		Authors: authors2,
	}

	items := []Item{item1, item2}

	image := Image{
		Title: "XML.com",
		Link:  "http://www.xml.com",
		URL:   "http://xml.com/universal/images/xml_tiny.gif",
	}

	links := []string{"http://xml.com/pub"}

	want := &Feed{
		RSSType:     RSS1,
		Title:       "Channel Title",
		Links:       links,
		Description: "this is description.",
		Image:       image,
		PubDate:     "2003-12-13T18:30:02Z",
		Items:       items,
	}

	if !reflect.DeepEqual(feed, want) {
		t.Error("TestFeedRSS1 not match expect feed, ", feed, want)
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
	feed, err := client.Parse(testServer.URL)
	if feed != nil || err == nil {
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
	feed, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedRSS2 fail get RSS response")
	}

	// assemble expecting feed
	links1 := []string{"http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp"}
	item1 := Item{
		ID:      "http://liftoff.msfc.nasa.gov/2003/06/03.html#item573",
		Title:   "Star City",
		Links:   links1,
		Body:    `How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.`,
		PubDate: "Tue, 03 Jun 2003 09:39:21 GMT",
	}

	item2 := Item{
		ID:      "http://liftoff.msfc.nasa.gov/2003/05/30.html#item572",
		Body:    `this is <b>bold</b>`,
		PubDate: "Fri, 30 May 2003 11:06:42 GMT",
	}

	links3 := []string{"http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp"}
	item3 := Item{
		ID:      "http://liftoff.msfc.nasa.gov/2003/05/27.html#item571",
		Title:   "The Engine That Does More",
		Links:   links3,
		Body:    `Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly. The proposed VASIMR engine would do that.`,
		PubDate: "Tue, 27 May 2003 08:37:32 GMT",
	}

	links4 := []string{"http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp"}
	item4 := Item{
		ID:      "http://liftoff.msfc.nasa.gov/2003/05/20.html#item570",
		Title:   "Astronauts' Dirty Laundry",
		Links:   links4,
		Body:    `Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them. Instead, astronauts have other options.`,
		PubDate: "Tue, 20 May 2003 08:56:02 GMT",
	}

	items := []Item{item1, item2, item3, item4}

	links := []string{"http://liftoff.msfc.nasa.gov/"}
	want := &Feed{
		RSSType:     RSS2,
		Title:       "Liftoff News",
		Links:       links,
		Description: "Liftoff to Space Exploration.",
		PubDate:     "Tue, 10 Jun 2003 04:00:00 GMT",
		Updated:     "Tue, 10 Jun 2003 09:41:01 GMT",
		Items:       items,
	}

	if !reflect.DeepEqual(feed, want) {
		t.Error("TestFeedRSS2 not match expect feed, ", feed, want)
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
	feed, err := client.Parse(testServer.URL)
	if feed != nil || err == nil {
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
	feed, err := client.Parse(testServer.URL)
	if err != nil {
		t.Error("TestFeedAtom fail get RSS response")
	}

	// assemble expecting feed
	links1 := []string{"http://example.org/2003/12/13/atom03"}
	item1 := Item{
		ID:      "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:   "This is Atom 1.0",
		Links:   links1,
		Body:    "Some text.",
		Updated: "2003-12-13T18:30:02Z",
		PubDate: "2003-12-13T08:29:29-04:00",
	}

	links2 := []string{"http://example.org/2003/12/13/atom03"}
	item2 := Item{
		ID:      "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:   "Today is hot",
		Links:   links2,
		Body:    "Some text.",
		Updated: "2003-12-13T18:30:02Z",
		PubDate: "2018-12-13T08:29:29-04:00",
	}

	items := []Item{item1, item2}
	links := []string{"http://google.com", "facebook.com"}

	a := Author{
		Name:  "John Doe",
		Email: "JohnDoe@example.com",
	}
	authors := []Author{a}

	image := Image{
		URL: "logo.jpg",
	}

	want := &Feed{
		RSSType:     Atom,
		Title:       "Example Feed",
		Links:       links,
		Description: "This is Sub title",
		Image:       image,
		Updated:     "2003-12-13T18:30:02Z",
		Authors:     authors,
		CopyRight:   "2005 naoto0822",
		Items:       items,
	}

	if !reflect.DeepEqual(feed, want) {
		t.Error("TestFeedAtom not match expect feed, ", feed, want)
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
	feed, err := client.Parse(testServer.URL)

	if feed != nil || err == nil {
		t.Error("TestErrorFeedAtom not expected return")
	}
}
