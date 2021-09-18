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
package main

import (
	"io/ioutil"
	"log"
	"time"
	"log"
	
	ht "github.com/imthaghost/merryGoRound/pkg/http"
)

func main() {
	// Configure a tor client
	tor := ht.Tor{
		MaxTimeout:         20 * time.Second,
		MaxIdleConnections: 10,
	}

	// new instance of tor client
	torClient := tor.New()

	// check your IP with AWS
	res, _ := torClient.Get("https://checkip.amazonaws.com")
	
	body, _ := ioutil.ReadAll(res.Body)
	ip := string(body)
	
	log.Printf("IP: %s", ip)
	
}

```