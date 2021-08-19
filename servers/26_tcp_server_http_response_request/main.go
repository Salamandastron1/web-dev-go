package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

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

	// read request
	url := request(c)

	// write response
	response(c, url)
}

func request(c net.Conn) string {
	i := 0
	scanner := bufio.NewScanner(c)
	var url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if i == 1 {
			url = strings.Fields(ln)[1]
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return url
}

func response(c net.Conn, url string) {
	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Dis the URL: %v &nbsp</strong></body></html>`, url)

	fmt.Fprint(c, "HTTP/1.1 200 ok\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}
