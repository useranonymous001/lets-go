package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	Method        string
	Path          string
	Headers       map[string]string
	Body          io.ReadCloser
	Conn          net.Conn
	Close         bool
	ContentLength int64
	URL           *url.URL
}

// methods for req type
// req.Body  --> accessing fields
// eg: req.ReadBodyAsString(req.body)
func (r *Request) ReadBodyAsString() (string, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
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
	var contentLength int64 = 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		kv := strings.SplitN(strings.TrimSpace(line), ":", 2)

		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			headers[key] = value

			if strings.EqualFold(key, "Content-Length") {
				contentLength, _ = strconv.ParseInt(value, 10, 64)
			}
		}
	}
	// handle post request
	var body = io.NopCloser(strings.NewReader(""))
	if contentLength > 0 {
		bodyBuffer := make([]byte, contentLength)
		_, err := io.ReadFull(reader, bodyBuffer)
		if err != nil {
			return nil, err
		}
		// converting []byte type to io.Reader type
		body = io.NopCloser(bytes.NewReader(bodyBuffer))
	}

	u, _ := url.Parse(path)

	return &Request{
		Method:        method,
		Path:          path,
		Close:         false,
		Headers:       headers,
		ContentLength: contentLength,
		Conn:          conn,
		Body:          body,
		URL:           u,
	}, nil
}
