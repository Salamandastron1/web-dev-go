package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var responseCodes map[string]int = map[string]int{
	"GET":    200,
	"DELETE": 204,
	"PATCH":  206,
	"POST":   201,
}

type requestLine struct {
	method      string
	uri         string
	httpVersion string
}

func main() {
	li, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		c, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	var rl requestLine
	// read request
	rl.request(c)

	// write response
	rl.response(c)
}

func (rl *requestLine) request(c net.Conn) {
	i := 0
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			fields := strings.Fields(ln)
			rl.method = fields[0]
			rl.uri = fields[1]
			rl.httpVersion = fields[2]
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func (rl *requestLine) response(c net.Conn) {
	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Dis the URL: %v &nbsp</strong><p>Dis the Method: %s</p></body></html>`, rl.uri, rl.method)
	rc, ok := responseCodes[rl.method];
	if !ok {
		fmt.Fprintf(c, "Incorrect HTTP Method %s", rl.method)
		return
	}
	fmt.Fprintf(c, "HTTP/1.1 %v ok\r\n", rc)
	if rl.method == "GET" {
		fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(c, "Content-Type: text/html\r\n")
		fmt.Fprint(c, "\r\n")
		fmt.Fprint(c, body)
	}
}
