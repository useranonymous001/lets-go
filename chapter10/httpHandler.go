package main

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HelloWriter struct {
	Message string
}

func MainHttp() {
	handler := HelloWriter{Message: "Hello, this is newbie learning interface in go"}
	handleFunc(HelloWriter{Message: "Hwllo usser"})
	http.ListenAndServe(":8080", handler)
}

func (h HelloWriter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message: %s", h.Message)
}

func handleFunc(handle Handler) {
	fmt.Printf("%+v\n", handle)
	h := handle.(HelloWriter)
	fmt.Printf("%+v", h)
}
