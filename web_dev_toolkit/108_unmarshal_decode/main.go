package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}
type thumbnail struct {
	Url           string
	Height, Width int
}

func main() {
	// marshal and unmarshal assign to a variable in the Go runtime while doing data transformation
	// encode and decode write to target "over the wire" while doing data transformation
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println("Error Unmarshalling", err)
	}
	fmt.Println(data)
}
