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
	u := models.User{Name: "james bond", Gender: "male", Age: 32}
	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Marshaling error:", err.Error())
	}
	b := bytes.NewReader(uj)
	r, _ := http.NewRequest("DELETE", "http://localhost:8080/user/61886241923f817dafc7ebe1", b)
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
