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
	limit       string
	attribution string
	version     string
	licence     string
	query       string
	typef       string
	//# features TODO: Add substruct.
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
	fmt.Printf("%s", response)

	json.Unmarshal(response, &parse)
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
func Reverse(parameters *map[string]string) {

}

/**
 * CSV search&reverse.
 */
func Csv() {

}

func main() {
	p := make(map[string]string)
	p["q"] = "1 allée des Bergeronnettes"

	var js *JSON
	js = Search(&p)
	println("\n\n\n\n")
	fmt.Printf("%s", js.limit)
}
