<img src="https://raw.githubusercontent.com/maelsan/gouvaddress/master/logo/france.jpeg" alt="France" width="200">

[![Software License](https://img.shields.io/badge/licence-MIT-blue.svg)](LICENSE)
[![codebeat badge](https://codebeat.co/badges/e1e6cf3b-821b-43f8-9318-d69d8ffdf1a7)](https://codebeat.co/projects/github-com-maelsan-gouvaddress)
[![Go Report Card](https://goreportcard.com/badge/github.com/maelsan/gouvaddress)](https://goreportcard.com/report/github.com/maelsan/gouvaddress)
[![Build Status](https://travis-ci.org/maelsan/gouvaddress.svg?branch=master)](https://travis-ci.org/maelsan/gouvaddress)
[![GoDoc](https://godoc.org/github.com/maelsan/gouvaddress?status.svg)](https://godoc.org/github.com/maelsan/gouvaddress)

## INSTALLATION
Just use `go get github.com/maelsan/gouvaddress` and import this package:

```go
import "github.com/maelsan/gouvaddress"
```

## DOCUMENTATION
I only implemented `/search/` & `/reverse/` because CSV format is too old, deprecated and so awful that I cannot imagine use it with a modern API. You must define a map of your parameters, like this:

```go
parameters := map[string]string{
		"q":   "8 bd du port",
		"lat": "48.357",
		"lon": "2.37",
	}
```

And call the function which correspond to the API endpoint.

```go
result := gouvaddress.Search(&parameters)
```

```go
result := gouvaddress.Reverse(&parameters)
```

The returns value is a struct named `JSON`, here is her composition:

```go
type JSON struct {
	Limit       int    `json:"limit"`
	Attribution string `json:"attribution"`
	Version     string `json:"version"`
	Licence     string `json:"licence"`
	Query       string `json:"query"`
	Type        string `json:"type"`
	Features    []struct {
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Citycode    string  `json:"citycode"`
			Postcode    string  `json:"postcode"`
			Name        string  `json:"name"`
			Housenumber string  `json:"housenumber"`
			Type        string  `json:"type"`
			Context     string  `json:"context"`
			Score       float64 `json:"score"`
			Label       string  `json:"label"`
			City        string  `json:"city"`
			ID          string  `json:"id"`
			Street      string  `json:"street"`
			Distance    int     `json:"distance"`
		} `json:"properties"`
		Type string `json:"type"`
	} `json:"features"`
}
```

Voilà! You can access to your result as a standard struct, for example:

```go
// Be careful with types, all objects in gouvaddress haven't the same type.
fmt.Printf("%s", result.Query)
fmt.Printf("%d", result.Limit)

// Some objects are substructs or substruct-arrays (so you can use loop or others...).
fmt.Printf("%s", result.Features[0].Properties.City)
fmt.Printf("%d", result.Features[0].Properties.Distance)
```

## Testing
You can test this package with `go test github.com/maelsan/gouvaddress`. It just checks if a result is returned.

## Licence
The MIT License (MIT). Please see License File for more information.