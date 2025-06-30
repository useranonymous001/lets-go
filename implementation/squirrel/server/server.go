package server

/*
Implementing Features of ServeMux:
	- ServeMux is an HTTP request multiplexer.
	- It matches the url of each incoming request against a list of registered
	patterns and calls the handler for the pattern that closely matches the URL
*/

import (
	"log"
	"net"
	"squirrel/core"
)

type SqurlMux struct {
	routes map[string]core.HandlerFunc
}

func SpawnServer() *SqurlMux {
	return &SqurlMux{
		routes: map[string]core.HandlerFunc{},
	}
}

func (sm *SqurlMux) Get(path string, handler core.HandlerFunc) {
	sm.routes["GET "+path] = handler
}

func (sm *SqurlMux) Post(path string, handler core.HandlerFunc) {
	sm.routes["POST "+path] = handler
}

func (sm *SqurlMux) Listen(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error while spawning server\n", err)
	}
	log.Println("Listening at", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error establishing connection", err)
		}

		// parse the request incoming to the connection: conn
		go func(conn net.Conn) {
			defer conn.Close()
			req, err := core.ParseRequest(conn)
			if err != nil {
				log.Fatal("Error parsing the request", err)
			}
			res := core.NewReponse(&conn)

			key := req.Method + " " + req.Path

			handlers, ok := sm.routes[key]

			if !ok {
				res.SetHeader("X-Squirrel", "Squirrel Data")
				res.SetStatus(404)
				res.Write("404 Not Found")
				res.Send()
			}

			handlers(req, res)
			res.Send()
		}(conn)
	}

}

// func ListenToServer(addr string) {

// 	ln, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Listening at", addr)

// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Fatal("failed to accept connection:", err)
// 		}
// 		go handleConnection(conn)
// 	}

// }

// func handleConnection(conn net.Conn) {

// 	defer conn.Close()

// 	req, err := core.ParseRequest(conn)

// 	if err != nil {
// 		log.Fatal("Error Parsing Request, ", err)
// 		return
// 	}

// 	core.ServeRequest(req)

// }
