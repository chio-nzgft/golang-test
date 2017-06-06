package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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


// mian 
func main() {
 // get & close web connect
	url := "http://localhost:1337"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

 // get body info
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

 // json.Unmarshal to Payload map
	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

  // print out Payload map
	fmt.Println(p.Stuff.Fruit, "\n", p.Stuff.Veggies)
}
