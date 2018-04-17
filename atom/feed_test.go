package atom

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	author1 := Author{
		Name:  "mike",
		Email: "hoge@gmail.com",
	}
	author2 := Author{
		Name:  "james",
		Email: "james@gmail.com",
	}
	authors := []Author{author1, author2}

	link1 := Link{
		Href: "https://facebook.com",
		Type: "hoge",
	}
	link2 := Link{
		Href: "https://google.com",
		Rel:  "foo",
	}
	links := []Link{link1, link2}

	e1 := Entry{
		ID:        "abcdefg",
		Title:     "entry1",
		Links:     links,
		Published: "2000-1-1",
		Updated:   "1900-1-1",
		Authors:   authors,
		Summary:   "this is summary",
	}

	e2 := Entry{
		ID:        "hogeid",
		Title:     "entry2",
		Links:     links,
		Published: "2000-1-1",
		Updated:   "1900-1-1",
		Authors:   authors,
		Content:   "this is content",
	}
	entries := []Entry{e1, e2}

	feed := Feed{
		ID:       "thisisID",
		Title:    "title",
		SubTitle: "sub title",
		Links:    links,
		Updated:  "2000-1-1",
		Authors:  authors,
		Icon:     "icon.png",
		Logo:     "logo.jpeg",
		Rights:   "naoto inc",
		Entries:  entries,
	}

	bytes, err := feed.ToJSON()
	ret := string(bytes)
	expect := `{"title":"title","links":["https://facebook.com","https://google.com"],"description":"sub title","updated":"2000-1-1","authors":[{"name":"mike","email":"hoge@gmail.com"},{"name":"james","email":"james@gmail.com"}],"image":{"url":"logo.jpeg"},"copyright":"naoto inc","categories":null,"items":[{"id":"abcdefg","title":"entry1","links":["https://facebook.com","https://google.com"],"description":"this is summary","content":"","pubdate":"2000-1-1","updated":"1900-1-1","authors":[{"name":"mike","email":"hoge@gmail.com"},{"name":"james","email":"james@gmail.com"}],"categories":null},{"id":"hogeid","title":"entry2","links":["https://facebook.com","https://google.com"],"description":"","content":"this is content","pubdate":"2000-1-1","updated":"1900-1-1","authors":[{"name":"mike","email":"hoge@gmail.com"},{"name":"james","email":"james@gmail.com"}],"categories":null}]}`

	if err != nil || ret != expect {
		t.Error("TestToJSON not match expected marshal string")
	}
}
