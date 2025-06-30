package core

import (
	"fmt"
	"net"
)

type Response struct {
	conn        net.Conn
	headers     map[string]string
	body        string
	contentType string
	statusCode  int
}

var (
	statusText = map[int]string{
		200: "OK",
		404: "Not Found",
		302: "Redirect",
		500: "Internal Server Error",
		// TODO: add more status texts as needed
	}
)

func NewReponse(conn *net.Conn) *Response {
	return &Response{
		conn:        *conn,
		statusCode:  200,
		contentType: "text/plain",
		headers:     map[string]string{},
	}
}

// methods to set headers fields
func (r *Response) SetHeader(key, value string) {
	r.headers[key] = value
}

func (r *Response) SetStatus(status int) {
	r.statusCode = status
}

func (r *Response) Write(body string) {
	r.body = body
}

func (r *Response) Send() {

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.statusCode, statusText[r.statusCode])
	contentLength := fmt.Sprintf("Content-Length: %d", len(r.body))
	contentType := fmt.Sprintf("Content-Type: %s", r.contentType)

	r.conn.Write([]byte(statusLine))
	r.conn.Write([]byte(contentType))
	r.conn.Write([]byte(contentLength))

	for k, v := range r.headers {
		r.conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}

	r.conn.Write([]byte("\r\n")) // end of headers

	// need to end with /r/n before writing the body
	r.conn.Write([]byte(r.body))

}
