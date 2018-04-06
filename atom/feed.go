package atom

import (
	"encoding/json"
	"html/template"

	// implement interfaces.Mappable
	_ "github.com/naoto0822/gss/interfaces"
)

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
	Summary      template.HTML `xml:"summary"`
	Contributors []Contributor `xml:"contributor"`
	Content      template.HTML `xml:"content"`
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

// ToJSON implmented interfaces.Mappable
// convert to map gss.Feed
func (f Feed) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}

// MarshalJSON assemble gss.Feed
func (f Feed) MarshalJSON() ([]byte, error) {
	type image struct {
		URL string `json:"url"`
	}
	i := image{URL: f.Logo}

	var links []Link
	if len(f.Links) > 0 {
		links = f.Links
	}

	var authors []Author
	if len(f.Authors) > 0 {
		authors = f.Authors
	}

	gf := &struct {
		Title       string   `json:"title"`
		Links       []Link   `json:"links"`
		Description string   `json:"description"`
		Updated     string   `json:"updated"`
		Authors     []Author `json:"authors"`
		Image       image    `json:"image"`
		CopyRight   string   `json:"copyright"`
		Items       []Entry  `json:"items"`
	}{
		Title:       f.Title,
		Links:       links,
		Description: f.SubTitle,
		Updated:     f.Updated,
		Authors:     authors,
		Image:       i,
		CopyRight:   f.Rights,
		Items:       f.Entries,
	}
	return json.Marshal(gf)
}

// MarshalJSON assemble gss Link
func (l Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.Href)
}

// MarshalJSON assemble gss.Author
func (a Author) MarshalJSON() ([]byte, error) {
	ga := &struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  a.Name,
		Email: a.Email,
	}
	return json.Marshal(ga)
}

// MarshalJSON assemble gss.Item
func (e Entry) MarshalJSON() ([]byte, error) {
	var body template.HTML
	if e.Summary != "" {
		body = e.Summary
	}
	if e.Content != "" {
		body = e.Content
	}

	var links []Link
	if len(e.Links) > 0 {
		links = e.Links
	}

	var authors []Author
	if len(e.Authors) > 0 {
		authors = e.Authors
	}

	gi := &struct {
		ID      string        `json:"id"`
		Title   string        `json:"title"`
		Links   []Link        `json:"links"`
		Body    template.HTML `json:"body"`
		PubDate string        `json:"pubdate"`
		Updated string        `json:"updated"`
		Authors []Author      `json:"authors"`
	}{
		ID:      e.ID,
		Title:   e.Title,
		Links:   links,
		Body:    body,
		PubDate: e.Published,
		Updated: e.Updated,
		Authors: authors,
	}
	return json.Marshal(gi)
}
