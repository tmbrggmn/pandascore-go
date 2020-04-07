# pandascore-go [![master](https://github.com/tmbrggmn/pandascore-go/workflows/master/badge.svg?branch=master)](https://github.com/tmbrggmn/pandascore-go/actions?query=workflow%3Amaster) [![codecov](https://codecov.io/gh/tmbrggmn/pandascore-go/branch/master/graph/badge.svg)](https://codecov.io/gh/tmbrggmn/pandascore-go) [![Go Report Card](https://goreportcard.com/badge/github.com/tmbrggmn/pandascore-go)](https://goreportcard.com/report/github.com/tmbrggmn/pandascore-go) [![GoDoc](https://godoc.org/github.com/tmbrggmn/pandascore-go?status.svg)](https://pkg.go.dev/github.com/tmbrggmn/pandascore-go)

[PandaScore](https://pandascore.co) client for Go.

## Status

:warning: This is still a new project and does *not* completely abstract all PandaScore APIs or games (yet). It also
has a couple of open issues (see below).

### Points of attention/improvement

 * Getting **all pages** from the PandaScore API has been implemented with by unmarshalling the results from all
 requests into a single array of `map[string]interface{}` to avoid having to use `reflect` to merge structs (see 
 [execution.go](execution.go). Not sure if this is the best way to tackle this problem :man_shrugging: 

## Getting started

### Installation

```
go get github.com/tmbrggmn/pandascore-go
```

### Basic usage

TODO

### Basic examples

TODO

## Usage

TODO

## Integration tests

In order to run the [integration tests](integration_test.go) you'll need to do 2 things:
 1. Add your on `.env` file to the root of the project 
 2. Enable integration tests by enabling the `integration` build flag

To set the **PandaScore access token**, simple create a new .env file in the root
of the project and add your [PandaScore access token](https://pandascore.co/settings) in the `PANDASCORE_ACCESS_TOKEN` 
variable. For example:

```dotenv
PANDASCORE_ACCESS_TOKEN=my_pandascore_access_token
```

Integration tests are marked with the `integration` build flag, so to run them
we need to **enable integration tests** when running the test: `go test -tags=integration` 
