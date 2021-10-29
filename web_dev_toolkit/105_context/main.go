package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type userID string
type fname string

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, userID("userID"), 777)
	ctx = context.WithValue(ctx, fname("fname"), "Bond")

	results := dbAccess(ctx)
	fmt.Fprintln(w, results)
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) int {
	return ctx.Value(userID("userID")).(int)
}
