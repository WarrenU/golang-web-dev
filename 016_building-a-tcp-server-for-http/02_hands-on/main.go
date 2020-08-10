package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><metacharset="UTF-8"><title></title></head><body>Hello</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
