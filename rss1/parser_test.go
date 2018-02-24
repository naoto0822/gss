package rss1

import (
	"testing"
)

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("TestNewParser not expected nil")
	}
}

func TestParseRSS1(t *testing.T) {
	t.Log("nothing test")
}
