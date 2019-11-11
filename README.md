# go-lightorama [![Go Report Card](https://goreportcard.com/badge/github.com/Cryptkeeper/go-lightorama)](https://goreportcard.com/report/github.com/Cryptkeeper/go-lightorama) [![GoDoc](https://godoc.org/github.com/Cryptkeeper/go-lightorama?status.svg)](https://godoc.org/github.com/Cryptkeeper/go-lightorama)
A Go library for controlling [Light-O-Rama (LOR) AC units](http://www1.lightorama.com/pro-ac-light-controllers/) such as the LOR160X series. 

This library is designed as a black box for higher level applications to control LOR units and is not a functional application of its own. 

## Usage
### Installation
Install using `go get github.com/Cryptkeeper/go-lightorama`

### Example Usage
See [example/example.go](example/example.go)

## Compatibility
`go-lightorama` re-implements the reverse engineered LOR protocol as documented at [github.com/Cryptkeeper/lightorama-protocol](https://github.com/Cryptkeeper/lightorama-protocol). As such this implementation *will* be feature incomplete and break with updates. Use with caution.

This library has only been tested with a `LOR1602WG3` unit.

See [TODO](TODO) for a list of absent, but planned features.

