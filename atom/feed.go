package atom

import ()

// cf. https://tools.ietf.org/html/rfc4287

// Feed atom feed
type Feed struct {
	ID       string   `xml:"id"`
	Title    string   `xml:"title"`
	SubTitle string   `xml:"subtitle"`
	Links    []Link   `xml:"link"`
	Updated  string   `xml:"updated"`
	Authors  []Author `xml:"author"`
	Icon     string   `xml:"icon"`
	Logo     string   `xml:"logo"`
	Rights   string   `xml:"rights"`
	Entries  []Entry  `xml:"entry"`
}

// Entry atom entry
type Entry struct {
	ID           string        `xml:"id"`
	Title        string        `xml:"title"`
	Links        []Link        `xml:"link"`
	Updated      string        `xml:"updated"`
	Published    string        `xml:"published"`
	Authors      []Author      `xml:"author"`
	Summary      string        `xml:"summary"`
	Contributors []Contributor `xml:"contributor"`
}

// Link atom link
type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

// Author atom author
type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	URI   string `xml:"uri"`
}

// Contributor atom contributor
type Contributor struct {
	Name string `xml:"name"`
}
