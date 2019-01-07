# trilaterate [![GoDoc](https://godoc.org/github.com/wilhelmstoll/trilaterate?status.svg)](https://godoc.org/github.com/wilhelmstoll/trilaterate) [![Build Status](https://travis-ci.org/wilhelmstoll/trilaterate.svg?branch=master)](https://travis-ci.org/wilhelmstoll/trilaterate)

Pure go implementation of trilaterate function, transcribed in part from Python originals by StackExchange user wwnick from post (http://gis.stackexchange.com/a/415/41129).


## Installation

```
go get -u github.com/wilhelmstoll/trilaterate
```

## Example of usage

```go
package main

import (
	"fmt"

	"github.com/wilhelmstoll/trilaterate"
)

func main() {
	b1 := trilaterate.Beacon{Lat: 35.000000, Lon: -120.000000, Dist: 189.419265289145}
	b2 := trilaterate.Beacon{Lat: 35.000005, Lon: -120.000010, Dist: 189.420325082156}
	b3 := trilaterate.Beacon{Lat: 35.000000, Lon: -120.000020, Dist: 189.420689733286}

	lat, lon := trilaterate.Solve(b1, b2, b3)

	latlon := fmt.Sprintf("%g,%g", lat, lon)
	fmt.Println(latlon)
}
```

## Reference

https://godoc.org/github.com/wilhelmstoll/trilaterate
