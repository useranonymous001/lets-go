package core

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
)

type Request struct {
	Method        string
	Path          string
	Headers       map[string]string
	Body          string
	Conn          net.Conn
	Close         bool
	ContentLength int64
	URL           *url.URL
}

func ParseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n') // reads the single line
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Fields(strings.TrimSpace(line))

	method, path := parts[0], parts[1]

	if method == "" {
		return nil, fmt.Errorf("INVALID METHOD")
	}

	headers := make(map[string]string)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		fmt.Println(line)
		kv := strings.SplitN(strings.TrimSpace(line), ":", 2)

		if len(kv) == 2 {
			headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return &Request{
		Method:        method,
		Path:          path,
		ContentLength: 0,
		Close:         false,
		Headers:       headers,
		Conn:          conn,
	}, nil
}
