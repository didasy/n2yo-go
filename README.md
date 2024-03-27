# n2yo-go
[![Build Status](https://travis-ci.org/JesusIslam/n2yo-go.svg?branch=master)](https://travis-ci.org/JesusIslam/n2yo-go)
[![codecov](https://codecov.io/gh/JesusIslam/n2yo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/JesusIslam/n2yo-go)
[![GoDoc](https://godoc.org/github.com/didasy/n2yo-go?status.svg)](https://godoc.org/github.com/didasy/n2yo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/didasy/n2yo-go)](https://goreportcard.com/report/github.com/didasy/n2yo-go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FJesusIslam%2Fn2yo-go.svg?type=small)](https://app.fossa.io/projects/git%2Bgithub.com%2FJesusIslam%2Fn2yo-go?ref=badge_small)

Unofficial [n2yo](https://n2yo.com) Golang client.

# Usage
Download and install it:

`go get github.com/didasy/n2yo-go`

Import it in your code:

`import "github.com/didasy/n2yo-go`

Simplest default example:

```
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/didasy/n2yo-go"
)

func main() {
	apiKey := os.Getenv("N2YO_API_KEY")
	client := n2yo.New(apiKey)
	issNORADID := 25544
	lat := -6.200000
	long := 106.816666
	alt := 5.0
	sec := 1
	days := 1
	minVisibility := 1
	minElevation := 15

	// get ISS TLE
	res, err := client.GetTLE(issNORADID)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("TLE:", string(data))

	// get ISS position
	res, err = client.GetPositions(issNORADID, lat, long, alt, sec)
	if err != nil {
		panic(err)
	}
	data, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Positions:", string(data))

	// get ISS visual passes
	res, err = client.GetVisualPasses(issNORADID, lat, long, alt, days, minVisibility)
	if err != nil {
		panic(err)
	}
	data, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Visual Passes:", string(data))

	// get ISS radio passes
	res, err = client.GetRadioPasses(issNORADID, lat, long, alt, days, minElevation)
	if err != nil {
		panic(err)
	}
	data, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Radio Passes:", string(data))
}

```