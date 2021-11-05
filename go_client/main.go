package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"web-dev-go/go_and_mongodb/organizing_code_into_packages/models"
)

func main() {
	u := models.User{Name: "james bond", Gender: "male", Age: 32, ID: "007"}
	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Marshaling error:", err.Error())
	}
	b := bytes.NewReader(uj)
	resp, err := http.Post("http://localhost:8080/user", "application/json", b)
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
