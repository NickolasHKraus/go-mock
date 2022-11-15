# Go Mock

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/nickolashkraus/go-mock/blob/master/LICENSE)

Examples for mocking Go code.

This repository uses [gomock](https://github.com/golang/mock), a mocking framework for the Go programming language.

## Philosophy

To understand testing and mocking in Go, you must understand that the struct you wish to mock must implement an interface, such that a mock struct can be used in its place when executed.

## Installation

Install the `mockgen` tool:

```
go install github.com/golang/mock/mockgen@latest
```

## Running `mockgen`

The `mockgen` command is used to generate source code for a mock struct given a Go source file containing interfaces to be mocked.

For a full list of flag, run `mockgen --help`.

To generate source code, run:

```
mockgen -package main -source main.go -destination main_mock.go
```

# TODO: time
var Now = time.Now

