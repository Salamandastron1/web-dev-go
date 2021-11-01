package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	Meow    string  `json:"Postal"`
	Moo     float64 `json:"Latitude"`
	Quack   float64 `json:"Longitude"`
	Honk    string  `json:"Address"`
	Neigh   string  `json:"City"`
	Bark    string  `json:"State"`
	Roar    string  `json:"Zip"`
	Screech string  `json:"Country"`
}

type model struct {
	State    bool
	Pictures []string
}

type cities []city

func main() {
	var data cities

	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling", err)
	}
	fmt.Println(data)

	fmt.Println(data[1].Bark)

	m := model{}
	fmt.Println("Before marshal, these are the zero values\n", m)
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error marshalling:", err)
	}

	fmt.Println("After JSON marshal these are the zero values:\n", string(bs))
}
