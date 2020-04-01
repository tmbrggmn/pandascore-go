# pandascore-go ![master](https://github.com/tmbrggmn/pandascore-go/workflows/master/badge.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/github/tmbrggmn/pandascore-go/badge.svg?branch=master)](https://coveralls.io/github/tmbrggmn/pandascore-go?branch=master) [![GoDoc](https://godoc.org/github.com/tmbrggmn/pandascore-go?status.svg)](https://pkg.go.dev/github.com/tmbrggmn/pandascore-go) [![Go Report Card](https://goreportcard.com/badge/github.com/tmbrggmn/pandascore-go)](https://goreportcard.com/report/github.com/tmbrggmn/pandascore-go)

[PandaScore](https://pandascore.co) client for Go.

## Getting started

### Installation

TODO

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