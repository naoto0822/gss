package main

import (
	"fmt"
	"io/ioutil"

	"github.com/naoto0822/go-rss/atom"
)

func main() {
	fmt.Println("hello go-rss!")

	path := "../testdata/atom_1.0.xml"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("failure read xml file")
	}

	parser := atom.Parser{}
	feed, err := parser.Parse(data)

	if err != nil {
		fmt.Errorf("failure parse atom xml")
	}

	fmt.Println("feed:", feed)
	fmt.Println("feed.Title:", feed.ID)
}
