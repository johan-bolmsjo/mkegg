# mkegg

`mkegg` is a really simple tool to embed arbitrary data into Go executable
files. It's to be used together with the `//go:generate` statement in Go source
files.

There are similar tools but this one is much simpler and only solves my
immediate needs. There is no registry of assets. Instead an exported function is
provided per asset that returns a `bytes.Reader` interface to read its data.

## Installation

    go install

## Usage

See `internal/eggs/...` for how to embed data and `mkegg_test.go` for how to
access it. Run `go generate ./...` once `mkegg` has been installed to update
source files with embedded data.
