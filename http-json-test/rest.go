package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// set up Payload struct
// Data include Fruits , Vegetables
// set as map
type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int
// end set Payload struct 

// for get get json to push http
func serveRest(w http.ResponseWriter, r *http.Request) {
 // use getJsonResponse() to get json
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
        
// push to response
	fmt.Fprintf(w, string(response))
}

// main
func main() {
	// Http HandleFunc from serveRest func
	http.HandleFunc("/", serveRest)
	// publish http
	http.ListenAndServe("localhost:1337", nil)
}

//Json data info return []byte, error
func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
