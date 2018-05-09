# onwater.io golang API wrapper

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/andy/onwater) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/andy/onwater/master/LICENSE)

[Onwater.io](https://onwater.io) is an API for determining if a point (latitude and longitude) is on water or land. This library is a small wrapper around for working with that API.

## Installation

```bash
go get github.com/andy/onwater
```

## Example

```go
package main

import (
  "context"
  "os"

  "github.com/andy/onwater"
)

func main() {
  client := onwater.New(os.Getenv("ONWATER_API_KEY")) // "" for an unpaid API Key with rate limits

  isWater, err := client.OnWater(context.Background(), 55.753675, 37.621339)
  if err != nil {
    // handle error which is most likely a http error or reply status code was not 200
  }

  if isWater {
    // lat/lng is on water
  } else {
    // lat/lng is on land
  }
}
```

> Note that there is also `client.OnLand` which is just an opposite for `client.OnWater`
