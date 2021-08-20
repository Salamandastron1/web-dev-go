package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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
	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Dis the URL: %v &nbsp</strong><h2>Dis the Method: %s</h2></body></html>`, rl.uri, rl.method)

	fmt.Fprint(c, "HTTP/1.1 200 ok\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}
