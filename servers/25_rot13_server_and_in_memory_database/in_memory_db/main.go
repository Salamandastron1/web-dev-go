package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	// instructions

	io.WriteString(c, "\nIN-MEMORY DATABASE\n\n"+`
USE:
SET key value
GET key
DEL key

EXAMPLE:
SET fav chocolate
GET fav

`)
	data := make(map[string]string)
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(c, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(c, "EXPECTED VALUE")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(c, "INVALID COMMAND")
		}
	}
}
