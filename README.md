# Azure Functions Go Handler

![stability-wip](https://img.shields.io/badge/stability-wip-lightgrey.svg)

This library allows you to write simple Go functions targeting Azure Functions
runtime without worrying too much about the custom handler logic.

The repo contains the custom handler logic and utilities for creating powerfull
serverless Go functions.

For basic usage see the `examples` for example use.

> Note that the project is in a very early state and breaking changes are likely to occur.

## Getting Started

If you don't already have it, install [Go](https://go.dev/dl/).

Get the lib: `go get -u github.com/NitorCreations/azure-functions-go-handler`

Use the lib: `import "github.com/NitorCreations/azure-functions-go-handler"`

Checkout `examples` on how to organize your code.
For now the handler code `examples/handler.go` must be created manually,
but the creationg could be automated by a simple command line tool (WIP).

## Features TODO

- Documentation
- Better support for the HTTP trigger
- Better support for other types of triggers
- Handler code generation tool
- ...

## Lisence

This project is licensed under MIT.
