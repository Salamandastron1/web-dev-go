package main

import (
	"log"
	"net/http"
)

func main() {
	// use Lets Encrypt to get free certficates
	// go run "/usr/local/go/src/crypto/tls/generate_cert.go" --host=localhost
	// or use some package on "pkg.go.dev"

	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server"))
}
