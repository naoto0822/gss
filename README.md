# gss

[![Build Status](https://travis-ci.org/naoto0822/gss.svg?branch=master)](https://travis-ci.org/naoto0822/gss)
[![GoDoc](https://godoc.org/github.com/naoto0822/gss?status.svg)](https://godoc.org/github.com/naoto0822/gss)
[![Go Report Card](https://goreportcard.com/badge/github.com/naoto0822/gss)](https://goreportcard.com/report/github.com/naoto0822/gss)
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/naoto0822/gss/blob/master/LICENSE)

go + rss = gss.  
gss is client to parse RSS Feed (rdf, rss, atom...).

### support
- [Atom](https://tools.ietf.org/html/rfc4287)
- [RSS1.0](http://web.resource.org/rss/1.0/spec)
- [RSS2.0](https://cyber.harvard.edu/rss/rss.html)
- [Content module](http://web.resource.org/rss/1.0/modules/content/) only encode (rss2)
- [Dublin Core module](http://web.resource.org/rss/1.0/modules/dc/) (rss1, rss2)
- [Media module](http://www.rssboard.org/media-rss) only part element (rss2)

## Installing

```sh
$ go get github.com/naoto0822/gss/gss
```

## Usage

### Parse RSS feed

```go
import "github.com/naoto0822/gss/gss"
```

Construct a new gss client, then use the `Parse` on the client to parse RSS feed.  
`Parse` returned `gss.Feed`.

```go
url := "https://jp.techcrunch.com/feed/"
client := gss.NewClient()
feed, err := client.Parse(url)
```

## Model Preview

`gss.Feed` is original and universal feed.  
refer [GoDoc](https://godoc.org/github.com/naoto0822/gss/gss#Feed) for details.

> TODO: rewrite following table style...

```go
feed
feed.RSSType
feed.Title
feed.Links
feed.Description
feed.Image
feed.CopyRight
feed.PubDate
feed.Updated
feed.Authors
feed.Categories
feed.Items

image
image.Title
image.URL
image.Link
image.Width
image.Height

item
item.ID
item.Title
item.Links
item.Body
item.PubDate
item.Updated
item.Authors
item.Categories

author
author.Name
author.Email
```

## Features

- [☓] support [Content module](http://web.resource.org/rss/1.0/modules/content/) only encode
- [☓] support [Dublin Core module](http://web.resource.org/rss/1.0/modules/dc/)
- [☓] support [Media module](http://www.rssboard.org/media-rss) only part element

## License

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.
