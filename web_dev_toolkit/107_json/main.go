package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First           string
	Last            string
	PersonalEffects []string
}

func main() {
	// marshal and unmarshal assign to a variable in the Go runtime while doing data transformation
	// encode and decode write to target "over the wire" while doing data transformation
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>You suck. Bye</h1>	
	</body>
	</html>`
	w.Write([]byte(s))
}
func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println("Marshalling failed:", err)
	}

	w.Write(json)
}
func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoding failed:", err)
	}
}
