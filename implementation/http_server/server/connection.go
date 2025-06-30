package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type Response struct {
	version     string
	contentLen  int
	contentType string
}

var (
	ErrNotFound = errors.New("404 Not Found")
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// read incoming data...
	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	headerPart := strings.Split(strings.TrimSpace(line), " ")

	method := headerPart[0]
	path := headerPart[1]

	body := "Hello, Mother Father !!"

	switch {
	case method == "GET" && path == "/hello":
		res := createResponse(body)
		fullResponse := fmt.Sprintf("%s\r\nContent-Length:%d\r\nContent-Type:%s\r\n\r\n%s", res.version, res.contentLen, res.contentType, body)
		conn.Write([]byte(fullResponse))

	default:
		notFoundRes := fmt.Sprintf("%s\r\nContent-Length:%d\r\nContent-Type:%s\r\n\r\n%s", "HTTP/1.1 404 Not Found", 13, "text/plain", ErrNotFound)
		conn.Write([]byte(notFoundRes))
	}

}

func createResponse(body string) Response {

	res := Response{
		contentLen:  len(body),
		contentType: "plain/text",
		version:     "HTTP/1.1 200 OK",
	}
	return res
}
