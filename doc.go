/*
This package use French Governmental Address API which is like a database of addresses in France. Note that those address are not affilied to citizens and do not contain personal informations. This API returns addresses referenced in France. It's like Google API, but with more accurate informations of French addresses.

You can find all informations about this API here: https://adresse.data.gouv.fr/api/

So you have to define your parameters, like this:
  parameters := map[string]string{
    "q":   "8 bd du port",
    "lat": "48.357",
    "lon": "2.37",
  }

And call the function which correspond to the API endpoint.
  result := gouvaddress.Search(&parameters)

or:
  result := gouvaddress.Reverse(&parameters)

Some concrete examples:
  // Be careful with types, all objects in gouvaddress haven't the same type.
  fmt.Printf("%s", result.Query)
  fmt.Printf("%d", result.Limit)

  // Some objects are substructs or substruct-arrays (so you can use loop or others...).
  fmt.Printf("%s", result.Features[0].Properties.City)
  fmt.Printf("%d", result.Features[0].Properties.Distance)

Values which are returned are accessible as follow:
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

For more informations https://github.com/maelsan/gouvaddress
*/
package gouvaddress
