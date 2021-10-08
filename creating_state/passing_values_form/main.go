package main

import (
	"io"
	"log"
	"net/http"
)

const (
	contentType  = "Content-Type"
	contentHTML  = "text/html; charset=UTF-8"
	contentPlain = "text/plain; charset=UTF-8"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	v := r.FormValue("v")
	w.Header().Set(contentType, contentHTML)

	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q" />
		<input type="submit" />
	</form>
	<br/>`+q)

	io.WriteString(w, `
	<form method="get">
		<input type="text" name="v" />
		<input type="submit" />
	</form>
	<br/>`+v)
}
