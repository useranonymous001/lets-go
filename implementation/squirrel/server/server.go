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
	"strings"
)

type route struct {
	method  string
	pattern string
	handler core.HandlerFunc
}

type SqurlMux struct {
	// routes map[string]core.HandlerFunc
	routes []route
}

func SpawnServer() *SqurlMux {
	// return &SqurlMux{
	// 	routes: map[string]core.HandlerFunc{},
	// }
	return &SqurlMux{}
}

func (sm *SqurlMux) Get(path string, handler core.HandlerFunc) {
	// sm.routes["GET "+path] = handler
	sm.routes = append(sm.routes, route{
		method:  "GET",
		pattern: path,
		handler: handler,
	})

}

func (sm *SqurlMux) Post(path string, handler core.HandlerFunc) {
	// sm.routes["POST "+path] = handler
	sm.routes = append(sm.routes, route{
		method:  "POST",
		pattern: path,
		handler: handler,
	})
}

func (sm *SqurlMux) Put(path string, handler core.HandlerFunc) {
	// sm.routes["PUT "+path] = handler
	sm.routes = append(sm.routes, route{
		method:  "PUT",
		pattern: path,
		handler: handler,
	})
}

func (sm *SqurlMux) Delete(path string, handler core.HandlerFunc) {
	// sm.routes["DELETE "+path] = handler

	sm.routes = append(sm.routes, route{
		method:  "DELETE",
		pattern: path,
		handler: handler,
	})

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

			params := map[string]string{}

			// checking the incoming request to the the each registered routes
			// comparing the route paths and setting the params if available (for dynamic routing)
			// if true, execute the corresponding handler for that routes registered
			// along with params{} if available
			for _, rt := range sm.routes {

				if rt.method != req.Method {
					continue
				}

				if matched := matchPattern(rt.pattern, req.Path, params); matched {
					req.Params = params
					rt.handler(req, res)
					res.Send()
					return
				}
			}
		}(conn)
	}

}

// returns true if the req.Path and rt.pattern matches
// comparing the incoming path with registered path

func matchPattern(pattern, path string, params map[string]string) bool {
	patternSegment := strings.Split(pattern, "/")
	pathSegment := strings.Split(path, "/")

	if len(pathSegment) != len(patternSegment) {
		return false
	}

	for i := 0; i < len(patternSegment); i++ {
		if strings.HasPrefix(patternSegment[i], ":") {
			paramName := strings.TrimPrefix(patternSegment[i], ":")
			params[paramName] = pathSegment[i]
			continue
		}
		if patternSegment[i] != pathSegment[i] {
			return false
		}
	}
	return true

}
