package main

import (
	"fmt"

	"github.com/naoto0822/gss/gss"
)

func main() {
	fmt.Println("hello go-rss!")

	url := "https://pitchfork.com/rss/news/"
	client := gss.NewClient()
	feed, err := client.Parse(url)

	if err != nil {
		fmt.Errorf("error:", err)
	}

	fmt.Printf("feed: %+v", feed)
}
