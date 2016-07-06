package gouvadress

import (
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
)

type NETAPI struct {
	domain string
	which  map[string]string
	from   string
	parameters url.Values
}

/**
 * Sets struct.
 */
func (newPI *NETAPI) netAPI(parameters *map[string]string, from string) {
	var setterApi NETAPI

	setterApi.domain = "api-adresse.data.gouv.fr"
	setterApi.which = map[string]string{
		"search":  "/search/",
		"reverse": "/reverse/",
		"csv":     "/csv/",
	}

	setterApi.addparameters(parameters)
	newPI = &setterApi
}

/**
 * Add parameters to NETAPI.
 */
func (setterApi *NETAPI) addparameters(parameters *map[string]string) {
	if len(*parameters) > 0 {
		for key, value := range *parameters {
			setterApi.parameters.Add(key, value)
		}
	}
}

/**
 * Decode response from API.
 */
func (setterApi *NETAPI) decode(method string) error {
	var (
		output interface{}
		r      *http.Response
		URI    string
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

	r = setterApi.execQuery(&method, &URI)

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(output)
}

/**
 * Executes a http query.
 */
func (setterApi *NETAPI) execQuery(method *string, URI *string) *http.Response {
	if *method == "GET" {
		r, e := http.Get(*URI)
		if e != nil {
			panic(e)
		}
		return r
	}
	r, e := http.PostForm(*URI, setterApi.parameters)
	if e != nil {
		panic(e)
	}
	return r
}

/**
 * Search API.
 */
func Search(parameters *map[string]string) error {
	var search *NETAPI

	search.netAPI(parameters, "Search")
	return search.decode("GET")
}

/**
 * Reverse API.
 */
func Reverse(parameters *map[string]string) error {
	var reverse *NETAPI

	reverse.netAPI(parameters, "Reverse")
	return reverse.decode("GET")
}

/**
 * CSV search&reverse.
 */
func Csv() {

}

func main() {
	p := make(map[string]string)
	p["q"] = "1 allée des Bergeronnettes"

	fmt.Printf(Search(&p))
}
