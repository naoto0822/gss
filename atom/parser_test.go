package atom

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseAtom(t *testing.T) {
	path := "../testdata/atom_1.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestParseAtom ioutil.ReadFile returned error:", err)
	}

	parser := Parser{}
	feed, err := parser.Parse(bytes)

	if err != nil {
		t.Error("TestParseAtom Parse returned error:", err)
	}

	if feed == nil {
		t.Error("TestParseAtom Parse returnd nil")
	}

	// following expected Struct
	var entryLinks []Link
	entryLink := Link{
		Href: "http://example.org/2003/12/13/atom03",
	}
	entryLinks = append(entryLinks, entryLink)

	var entryContributors []Contributor
	contributor := Contributor{Name: "Sam Ruby"}
	entryContributors = append(entryContributors, contributor)

	entry1 := Entry{
		ID:           "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:        "This is Atom 1.0",
		Links:        entryLinks,
		Updated:      "2003-12-13T18:30:02Z",
		Published:    "2003-12-13T08:29:29-04:00",
		Summary:      "Some text.",
		Contributors: entryContributors,
	}

	var entryLinks2 []Link
	entryLink2 := Link{
		Href: "http://example.org/2003/12/13/atom03",
	}
	entryLinks2 = append(entryLinks2, entryLink2)
	entry2 := Entry{
		ID:        "urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a",
		Title:     "Today is hot",
		Links:     entryLinks2,
		Updated:   "2003-12-13T18:30:02Z",
		Published: "2018-12-13T08:29:29-04:00",
		Summary:   "Some text.",
	}

	var entries []Entry
	entries = append(entries, entry1)
	entries = append(entries, entry2)

	var feedLinks []Link
	link1 := Link{
		Href: "http://google.com",
		Type: "application/atom+xml",
		Rel:  "alternate",
	}
	link2 := Link{
		Href: "facebook.com",
		Type: "application/atom+xml",
		Rel:  "self",
	}
	feedLinks = append(feedLinks, link1)
	feedLinks = append(feedLinks, link2)

	var authors []Author
	author := Author{
		Name:  "John Doe",
		Email: "JohnDoe@example.com",
		URI:   "http://example.com/~johndoe",
	}
	authors = append(authors, author)

	want := &Feed{
		ID:       "urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6",
		Title:    "Example Feed",
		SubTitle: "This is Sub title",
		Links:    feedLinks,
		Updated:  "2003-12-13T18:30:02Z",
		Authors:  authors,
		Icon:     "icon.jpg",
		Logo:     "logo.jpg",
		Rights:   "2005 naoto0822",
		Entries:  entries,
	}

	if !reflect.DeepEqual(feed, want) {
		t.Error("TestParseAtom Parse returned not match, ", feed, want)
	}

}

func TestParseErrorAtom(t *testing.T) {

}
