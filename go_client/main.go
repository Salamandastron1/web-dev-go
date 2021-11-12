package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// watch out for those zero val
	uj := `{"name": "james", "gender": "she/her", "age": 18}`
	b := bytes.NewReader([]byte(uj))
	r, err := http.NewRequest("POST", "http://localhost:8080/user", b)
	if err != nil {
		log.Fatal("Bad inputs:", err.Error())
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(body))
}
