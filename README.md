# goharbor-client
[![GitHub license](https://img.shields.io/github/license/zelat/goharbor-client.svg?style=flat-square)](https://github.com/zelat/goharbor-client/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zelat/goharbor-client?style=flat-square)](https://goreportcard.com/badge/github.com/zelat/goharbor-client)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/zelat/goharbor-client/)
[![Release](https://img.shields.io/github/release/zelat/goharbor-client.svg?style=flat-square)](https://github.com/zelat/goharbor-client/releases/latest)

[![Maintainability](https://api.codeclimate.com/v1/badges/a765bafaa29f6f8fdde7/maintainability)](https://codeclimate.com/github/zelat/goharbor-client/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a765bafaa29f6f8fdde7/test_coverage)](https://codeclimate.com/github/zelat/goharbor-client/test_coverage)
[![Actions Status](https://github.com/zelat/goharbor-client/workflows/Test/badge.svg)](https://github.com/zelat/goharbor-client/actions)

A Go client library enabling programs to perform CRUD operations on the [goharbor](https://github.com/goharbor/harbor) API.

This client library utilizes types generated by [go-swagger](https://github.com/go-swagger/go-swagger).

## Compatibility
This library includes separate clients supporting the **v1.10** & **v2.2** versions of goharbor. 

## Installation
Install the desired client library version using `go get`:

```shell script
# v1 Client
go get github.com/zelat/goharbor-client/apiv1
```

or

```shell script
# v2 Client
go get github.com/zelat/goharbor-client/apiv2
```

## Contributing
Before you make your changes, check to see if an [issue already exists](https://github.com/zelat/goharbor-client/issues) for the change you want to make.

When in doubt where to start when making changes to the client, please refer to the [Contribution guide](./CONTRIBUTING.md).

## Documentation
For more specific documentation, please refer to the [godoc](https://pkg.go.dev/github.com/zelat/goharbor-client/) of this library.
