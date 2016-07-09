package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type NETAPI struct {
	domain     string
	which      map[string]string
	from       string
	parameters url.Values
}

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
		} `json:"properties"`
		Type string `json:"type"`
	} `json:"features"`
}

/**
 * Sets struct.
 */
func netAPI(parameters *map[string]string, from string) *NETAPI {
	var setterApi NETAPI

	setterApi.from = from
	setterApi.domain = "http://api-adresse.data.gouv.fr"
	setterApi.which = map[string]string{
		"search":  "/search/",
		"reverse": "/reverse/",
		"csv":     "/csv/",
	}

	setterApi.addparameters(parameters)
	return &setterApi
}

/**
 * Add parameters to NETAPI.
 */
func (setterApi *NETAPI) addparameters(parameters *map[string]string) {
	if len(*parameters) > 0 {
		setterApi.parameters = make(map[string][]string)
		for key, value := range *parameters {
			setterApi.parameters.Add(key, value)
		}
	}
}

/**
 * Decode response from API.
 */
func (setterApi *NETAPI) decode(method string) *JSON {
	var (
		response []byte
		URI      string
		parse    JSON
	)
	URI = setterApi.domain

	switch setterApi.from {
	case "Search":
		URI += setterApi.which["search"]
	case "Reverse":
		URI += setterApi.which["reverse"]
	case "Csv":
		URI += setterApi.which["csv"]
	}

	URI += "?"
	URI += setterApi.parameters.Encode()

	response = setterApi.execQuery(&method, &URI)

	r := json.Unmarshal(response, &parse)
	if r != nil {
		panic(r)
	}

	return &parse
}

/**
 * Executes a http query.
 * TODO: Do with go routine
 * and see if we can use closure
 */
func (setterApi *NETAPI) execQuery(method *string, URI *string) []byte {
	if *method == "GET" {
		r, e := http.Get(*URI)
		if e != nil {
			panic(e)
		}

		body, _ := ioutil.ReadAll(r.Body)
		return body
	}
	r, e := http.PostForm(*URI, setterApi.parameters)
	if e != nil {
		panic(e)
	}

	body, _ := ioutil.ReadAll(r.Body)
	return body
}

/**
 * Search API.
 */
func Search(parameters *map[string]string) *JSON {
	var search *NETAPI

	search = netAPI(parameters, "Search")
	return search.decode("GET")
}

/**
 * Reverse API.
 */
func Reverse(parameters *map[string]string) *JSON {
	var reverse *NETAPI

	reverse = netAPI(parameters, "Reverse")
	return reverse.decode("GET")
}

func main() {
	pp := make(map[string]string)
	pp["lon"] = "2.37"
	pp["lat"] = "48.357"

	var jp *JSON
	jp = Reverse(&pp)
	fmt.Printf("%#v", jp)
}
