package gouvaddress

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NETAPI Struct to define URI API.
type NETAPI struct {
	domain     string
	which      map[string]string
	from       string
	parameters url.Values
}

// JSON Struct to store JSON response.
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

// netAPI Set struct before send request.
func netAPI(parameters *map[string]string, from string) *NETAPI {
	var setterAPI NETAPI

	setterAPI.from = from
	setterAPI.domain = "http://api-adresse.data.gouv.fr"
	setterAPI.which = map[string]string{
		"search":  "/search/",
		"reverse": "/reverse/",
	}

	setterAPI.addParameters(parameters)
	return &setterAPI
}

// addParameters Add parameters given from dev call.
func (setterAPI *NETAPI) addParameters(parameters *map[string]string) {
	if len(*parameters) > 0 {
		setterAPI.parameters = make(map[string][]string)
		for key, value := range *parameters {
			setterAPI.parameters.Add(key, value)
		}
	}
}

// decode Decode response from JSON.
func (setterAPI *NETAPI) decode(method string) *JSON {
	var (
		response []byte
		URI      string
		parse    JSON
	)
	URI = setterAPI.domain

	switch setterAPI.from {
	case "Search":
		URI += setterAPI.which["search"]
	case "Reverse":
		URI += setterAPI.which["reverse"]
	}

	URI += "?"
	URI += setterAPI.parameters.Encode()

	response = setterAPI.execQuery(&method, &URI)

	r := json.Unmarshal(response, &parse)
	if r != nil {
		panic(r)
	}

	return &parse
}

// execQuery Send a HTTP query.
func (setterAPI *NETAPI) execQuery(method *string, URI *string) []byte {
	if *method == "GET" {
		r, e := http.Get(*URI)
		if e != nil {
			panic(e)
		}

		body, _ := ioutil.ReadAll(r.Body)
		return body
	}
	r, e := http.PostForm(*URI, setterAPI.parameters)
	if e != nil {
		panic(e)
	}

	body, _ := ioutil.ReadAll(r.Body)
	return body
}

// Search function to use /search/ API
func Search(parameters *map[string]string) *JSON {
	var search *NETAPI

	search = netAPI(parameters, "Search")
	return search.decode("GET")
}

// Reverse function to use /reverse/ API
func Reverse(parameters *map[string]string) *JSON {
	var reverse *NETAPI

	reverse = netAPI(parameters, "Reverse")
	return reverse.decode("GET")
}
