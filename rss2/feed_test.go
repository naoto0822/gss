package rss2

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	guid := GUID{
		Value:       "0123456789",
		IsPermaLink: "link",
	}

	category1 := Category{
		Value: "sports",
	}
	category2 := Category{
		Value: "music",
	}
	categories := []Category{category1, category2}

	item1 := Item{
		GUID:        guid,
		Title:       "rss2 article",
		Link:        "https://facebook.com",
		Description: "<p>hahaha</p>",
		Content:     "<p>this is content</p>",
		Author:      "naoto",
		Categories:  categories,
		PubDate:     "this is publish date",
	}
	items := []Item{item1}

	image := Image{
		Title:  "hoge.png",
		URL:    "http://hoge.png",
		Link:   "https://twitter.com",
		Width:  300,
		Height: 500,
	}

	channel := Channel{
		Title:         "music feed",
		Link:          "https://yahoo.co.jp",
		Description:   "this is rss2 feed",
		Language:      "ja",
		CopyRight:     "Mr.hoge",
		WebMaster:     "Foo",
		PubDate:       "1900-1-1",
		LastBuildDate: "2000-1-1",
		Categories:    categories,
		Image:         image,
		Rating:        "5",
		Items:         items,
	}

	feed := Feed{
		Channel: channel,
	}

	bytes, err := feed.ToJSON()
	ret := string(bytes)
	expect := `{"title":"music feed","link":"https://yahoo.co.jp","description":"this is rss2 feed","image":{"title":"hoge.png","url":"http://hoge.png","link":"https://twitter.com","width":300,"height":500},"copyright":"Mr.hoge","pubdate":"1900-1-1","updated":"2000-1-1","categories":["sports","music"],"items":[{"id":"0123456789","title":"rss2 article","link":"https://facebook.com","description":"\u003cp\u003ehahaha\u003c/p\u003e","content":"\u003cp\u003ethis is content\u003c/p\u003e","pubdate":"this is publish date","authors":[{"name":"naoto","email":""}],"categories":["sports","music"],"enclosure":{"url":"","length":0,"type":""}}]}`

	if err != nil || ret != expect {
		t.Error("TestToJSON not match expected marshal string")
	}
}
