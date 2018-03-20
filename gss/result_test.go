package gss

import (
	"testing"
)

func TestResultIsSuccessful(t *testing.T) {
	result := Result{}
	result.isSuccessful = true

	if result.IsSuccessful() != true {
		t.Error("TestResultIsSuccessful not expected result")
	}
}

func TestResultIsRSS1(t *testing.T) {
	result := Result{}
	result.RSSType = RSS1

	if result.IsRSS1() != true {
		t.Error("TestResultIsRSS1 not expected result")
	}
}

func TestResultIsRSS2(t *testing.T) {
	result := Result{}
	result.RSSType = RSS2

	if result.IsRSS2() != true {
		t.Error("TestResultIsRSS2 not expected result")
	}
}

func TestResultIsAtom(t *testing.T) {
	result := Result{}
	result.RSSType = Atom

	if result.IsAtom() != true {
		t.Error("TestResultIsAtom not expected result")
	}
}
