# Azure Functions Go Handler

![stability-wip](https://img.shields.io/badge/stability-wip-lightgrey.svg)

> Note that the project is in a very early state and breaking changes are likely to occur.

This library allows you to write simple Go functions targeting Azure Functions
runtime without worrying too much about the custom handler logic.

The repo contains the custom handler logic and utilities for creating powerfull
serverless Go functions.

For basic usage see `examples`.

## Getting Started

If you don't already have it, install [Go](https://go.dev/dl/).

Start using:

- Get the lib: `go get -u github.com/NitorCreations/azure-functions-go-handler`
- Use the lib: `import "github.com/NitorCreations/azure-functions-go-handler"`

Use `gofunc` CLI to init a new project

- Install `gofunc` if needed (see [Cli Reference](#cli-reference))
- Create a new project dir and cd: `mkdir go-func-app && cd go-func-app`
- Initialize project structure `gofunc init`

## Usage

The library provides common utils to handle invocations by the Function Runtime and leave the function logic implementation to the developers.

> TODO describe handler function allowed signatures

## CLI Reference

The `gofunc` Command Line Interface provides utilities to help with Go Function Apps.

### Installation

```shell
go install -i github.com/NitorCreations/azure-functions-go-handler/cmd/gofunc
```

### Synopsis

```
Usage:

    gofunc <command> [parameters]

The commands are

    init             create a new Go Function App in the current directory
    generate [dir]   generate func handler code starting from optional [dir], defaults to current directory
    version          print version info and exit
    help             show this help
```

### Init Command

Creates a new Go Function App in the working directory with an example HttpTrigger that is ready to Go.

```shell
gofunc init
```

### Generate Command

Generates custom handler code from the available functions. Parameter `dir` is optional and defaults to current directory.

```
gofunc generate [dir]
```

Functions are detected by folders containing a `function.json` definition file.
It is expected that the entry point function can be found in the same directory, has a package name equal to the directory and is named `Handle`.

To use a different entry point function set the `entryPoint` property in `functions.json`.

To skip the function completely in generation set the `excluded` property to `true` in `functions.json`.

## Features TODO

- Documentation
- Release automation, simple tests
- Better support for the HTTP trigger
- Better support for other types of triggers
- ...

## Lisence

This project is licensed under MIT.
