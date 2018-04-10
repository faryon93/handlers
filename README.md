# Handlers
[![Documentation](https://godoc.org/github.com/faryon93/handlers?status.svg)](http://godoc.org/github.com/faryon93/handlers)
[![Go Report Card](https://goreportcard.com/badge/github.com/faryon93/handlers)](https://goreportcard.com/report/github.com/faryon93/handlers)
[![Last Release](https://img.shields.io/github/release/faryon93/handlers.svg)](https://github.com/faryon93/handlers/releases)

A set of frequently used Go http handlers and middleware functions.

    # download library
    $: go get github.com/faryon93/handlers

    # place on top of your go file
    import "github.com/faryon93/handlers"

## Content: Handlers

| Handler       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| NoRobots()    | Writes a robots.txt file, which disallows the access to everything. |
| Forbidded()   | Default 403 forbidden handler.                                      |

## Content: Adapters

Adapter functions can be chained with the real handler function or other adapter functions.

| Adapter           | Description                                                       |
| ----------------- | ----------------------------------------------------------------- |
| Keyed(reqKey)     | Restrict access to requests, having param "key" matching reqKey.  |
| Enabled(en)       | Denys access if en is false                                       |
| Benchmark         | Logs the execution time of every request using logrus             |
| Paged(limit)      | Paging: Parses skip and limit from query parameters               |