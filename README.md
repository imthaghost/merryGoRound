# merryGoRound

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/tempor1s/notiongo)
[![Test Status](https://github.com/google/go-github/workflows/tests/badge.svg)](https://github.com/google/go-github/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/google/go-github/branch/master/graph/badge.svg)](https://codecov.io/gh/google/go-github)


merryGoRound is a simple HTTP Client with rotating IPs via SmartProxy or Tor.


## Installation ##

merryGoRound is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/imthaghost/merrygoround/
```

will resolve and add the package to the current development module, along with its dependencies.

## Usage ##

```go
import  mgr "github.com/imthaghost/merrygoround"
```

Construct a new merryGoRound client, then use the various services on the client to
access different parts of the Notion API. For example:

```go
client := mgr.NewTorClient()
// check your IP with AWS
res, err := client.Get("https://checkip.amazonaws.com")
```

### Integration Tests ###

You can run integration tests from the `test` directory. See the integration tests [README](test/README.md).

