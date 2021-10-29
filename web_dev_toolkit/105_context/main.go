package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type userID string
type fname string

var uID userID = "userID"

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, uID, 777)
	ctx = context.WithValue(ctx, fname("fname"), "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, results)
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		// ridiculous long running task
		uid := ctx.Value(uID).(int)
		time.Sleep(10 * time.Second)

		//check to mae sure we're not running in vain
		// if ctx.Done()
		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}
