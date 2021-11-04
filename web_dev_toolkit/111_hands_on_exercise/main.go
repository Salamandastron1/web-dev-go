package main

import (
	"encoding/json"
	"fmt"
)

type code struct {
	Code string
	Num  int
}

func main() {
	rcvd := `[{"code":"StatusOK","num":200},{"code":"StatusFound","num":301},{"code":"StatusForbidden","num":403},{"code":"StatusTeapot","num":418}]`
	var data []code
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println("Failed to unmarshal:", err)
	}
	fmt.Println(data)
	for _, v := range data {
		fmt.Println(v.Code)
		fmt.Println(v.Num)
	}
}
