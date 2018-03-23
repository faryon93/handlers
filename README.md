# Handlers
[![Documentation](https://godoc.org/github.com/faryon93/handlers?status.svg)](http://godoc.org/github.com/faryon93/handlers)
[![Go Report Card](https://goreportcard.com/badge/github.com/faryon93/handlers)](https://goreportcard.com/report/github.com/faryon93/handlers)

A set of frequently used Go http handlers and middleware functions.

    # download library
    $: go get github.com/faryon93/handlers

    # place on top of your go file
    import "github.com/faryon93/handlers"

## Content: Handlers

| Handler       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| NoRobots()    | Writes a robots.txt file, which disallows the access to everything. |

## Content: Middleware

Middleware functions can be chained with the real handler function or other middle ware functions.

| Middleware        | Description                                                       |
| ----------------- | ----------------------------------------------------------------- |
| Keyed(reqKey)     | Restrict access to requests, having param "key" matching reqKey.  |
| Enabled(confKey)  | Denys access if the resolved bool value of confKey is false.      |
