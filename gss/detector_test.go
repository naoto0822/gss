package gss

import (
	"io/ioutil"
	"testing"
)

func TestDetectAtom(t *testing.T) {
	path := "../testdata/atom_1.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestDetectAtom ioutil.ReadFile returnd error:", err)
	}

	detector := NewDetector()

	root, err := detector.detect(bytes)
	if err != nil {
		t.Error("TestDetectAtom Detect not expected error:", err)
	}

	t.Log(root)
}

func TestDetectRSS1(t *testing.T) {
	path := "../testdata/rss_1.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestDetectRSS1 ioutil.ReadFile returned error:", err)
	}

	detector := NewDetector()
	root, err := detector.detect(bytes)
	if err != nil {
		t.Error("TestDetectRSS1 Detect not expected error:", err)
	}

	t.Log(root)
}

func TestDetectRSS2(t *testing.T) {
	path := "../testdata/rss_2.0.xml"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error("TestDetectRSS2 ioutil.ReadFile returned error:", err)
	}

	detector := NewDetector()
	root, err := detector.detect(bytes)
	if err != nil {
		t.Error("TestDetectRSS2 Detect not expected error:", err)
	}

	t.Log(root)
}
