package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Wow you wrote some code that responds to stuff with HTTP syntax")
}
func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
