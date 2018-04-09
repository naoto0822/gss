# gss

[![Build Status](https://travis-ci.org/naoto0822/gss.svg?branch=master)](https://travis-ci.org/naoto0822/gss)
[![GoDoc](https://godoc.org/github.com/naoto0822/gss?status.svg)](https://godoc.org/github.com/naoto0822/gss)
[![Go Report Card](https://goreportcard.com/badge/github.com/naoto0822/gss)](https://goreportcard.com/report/github.com/naoto0822/gss)
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/naoto0822/gss/blob/master/LICENSE)

rss pkg of go lang.

### support
- [Atom](https://tools.ietf.org/html/rfc4287)
- [RSS1.0](http://web.resource.org/rss/1.0/spec)
- [RSS2.0](https://cyber.harvard.edu/rss/rss.html)

## Installing

To parse rss

```sh
$ go get github.com/naoto0822/gss/gss
```

## Usage

### Parse RSS feed

```go
import "github.com/naoto0822/gss/gss"
```

Construct a new gss client, then use the `Parse` on the client to parse RSS feed. 

```go
url := "https://jp.techcrunch.com/feed/"
client := gss.NewClient()
feed, err := client.Parse(url)
```

### Generate Feed

> feature...

## Features

- generating feed

## License

This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE)
file.
