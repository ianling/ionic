# Ionic

[![Ion Channel Status](https://api.ionchannel.io/v1/report/getBadge?project_id=1166459a-15fe-420b-856f-874e612b08a6&branch=master)](http://console.ionchannel.io/)
![Build Status](https://github.com/ion-channel/ionic/workflows/Build/badge.svg)
[![Go Reportcard](https://goreportcard.com/badge/github.com/ion-channel/ionic)](https://goreportcard.com/report/github.com/ion-channel/ionic)
[![GoDoc](https://godoc.org/github.com/ion-channel/ionic?status.svg)](https://godoc.org/github.com/ion-channel/ionic)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/ion-channel/ionic/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/ion-channel/ionic.svg)](https://github.com/ion-channel/ionic/releases/latest)

A Go SDK for the Ion Channel API.

Some structs are automatically generated using [Ion Channel's GraphQL schema](https://github.com/ion-channel/graphql-schema)
and [gqlgen](https://github.com/99designs/gqlgen).

The schema can be retrieved by running `make get_schema`, which will download the version of the schema specified near the top of the Makefile.
Structs can be regenerated from the schema by running `make generate`.

# Requirements
Go Version 1.17 or higher

# Installation
```
go get github.com/ion-channel/ionic
```

# Versioning
The SDK will be versioned in accordance with [Semver 2.0.0](http://semver.org).  See the [releases](https://github.com/ion-channel/ionic/releases) section for the latest version.
Until version 1.0.0 the SDK is considered to be unstable.

# License
This SDK is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).  See [LICENSE.md](./LICENSE.md) for more information.
