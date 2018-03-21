package gss

import (
	"testing"
)

func TestGetParser(t *testing.T) {
	{
		rssType := RSS1
		p, err := getParser(rssType)
		if p == nil || err != nil {
			t.Error("TestGetParser failed get rss1 Parser")
		}
	}
	{
		rssType := RSS2
		p, err := getParser(rssType)
		if p == nil || err != nil {
			t.Error("TestGetParser failed get rss2 Parser")
		}
	}
	{
		rssType := Atom
		p, err := getParser(rssType)
		if p == nil || err != nil {
			t.Error("TestGetParser failed get atom Parser")
		}
	}
	{
		rssType := Unknown
		p, err := getParser(rssType)
		if p != nil || err == nil {
			t.Error("TestGetParser failed detect unknown type")
		}
	}
}
