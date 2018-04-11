package rss1

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	channel := Channel{
		Title:       "this is title",
		Link:        "https://google.com",
		Description: "this is description",
	}

	var items []Item
	item := Item{
		Title:       "this is item title",
		Link:        "https://yahoo.co.jp",
		Description: "this is description",
	}
	items = append(items, item)

	feed := Feed{
		Channel: channel,
		Items:   items,
	}

	bytes, err := feed.ToJSON()
	ret := string(bytes)
	expect := `{"title":"this is title","links":["https://google.com"],"description":"this is description","image":{"title":"","url":"","link":""},"pubdate":"","items":[{"title":"this is item title","links":["https://yahoo.co.jp"],"description":"this is description","pubdate":"","authors":null}]}`

	if err != nil || ret != expect {
		t.Error("TestToJSON not match ecpected marshal string")
	}
}
