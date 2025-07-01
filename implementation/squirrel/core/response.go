package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

type Response struct {
	conn    net.Conn
	headers map[string]string
	// body        string
	body        io.ReadCloser
	contentType string
	statusCode  int
	close       bool
}

var (
	statusText = map[int]string{
		200: "OK",
		404: "Not Found",
		302: "Redirect",
		500: "Internal Server Error",
		403: "Forbidden",
		401: "Bad Request",
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

// our body is io.ReadCloser that accepts Reader type
// NopCloser returns a [ReadCloser] with a no-op Close method wrapping the provided [Reader] r.
// strings.NewReader(string) converts our string to Reader type
func (r *Response) Write(body string) {
	r.body = io.NopCloser(strings.NewReader(body))
}

// helper method, in any case of Use a File or Stream
func (r *Response) SetBody(reader io.ReadCloser) {
	r.body = reader
}

/*
## Accpets body type as raw bytes as well
Works great for:
Binary responses
Images
Files
Compressed/gzipped data
*/
func (r *Response) WriteBytes(b []byte) {
	r.body = io.NopCloser(bytes.NewReader(b))
}

func (r *Response) Close() {
	r.close = true
}

func (r *Response) JSON(data interface{}) {
	// data need to be a a type struct(atleast)
	b, err := json.MarshalIndent(data, "", "")

	if err != nil {
		r.conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	// all headers are set in the send() methods
	r.SetHeader("Content-Type", "application/json")
	r.contentType = "application/json"

	// setting body
	r.WriteBytes(b)
}

// send
func (r *Response) Send() {
	// if body is empty, default empty body
	if r.body == nil {
		r.Write("")
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.statusCode, statusText[r.statusCode])
	contentType := fmt.Sprintf("Content-Type: %s\r\n", r.contentType)

	// A Buffer is a variable-sized buffer of bytes with [Buffer.Read] and [Buffer.Write] methods.
	// The zero value for Buffer is an empty buffer ready to use.
	var bodyBuf bytes.Buffer
	defer r.body.Close()
	_, err := io.Copy(&bodyBuf, r.body)
	if err != nil {
		r.conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	contentLength := fmt.Sprintf("Content-Length: %d\r\n", bodyBuf.Len())

	r.conn.Write([]byte(statusLine))
	r.conn.Write([]byte(contentType))
	r.conn.Write([]byte(contentLength))

	for k, v := range r.headers {
		r.conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}

	r.conn.Write([]byte("\r\n")) // end of headers

	// need to end with /r/n before writing the body
	r.conn.Write(bodyBuf.Bytes())
}
