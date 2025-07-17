package server

/*
Implementing Features of ServeMux:
	- ServeMux is an HTTP request multiplexer.
	- It matches the url of each incoming request against a list of registered
	patterns and calls the handler for the pattern that closely matches the URL
*/

/*

	Explanation:
	routeHandler = rt.middlewares[i](routeHandler)

	Our Middleware is something like this:

	--> type Middleware func(core.HandlerFunc) core.HandlerFunc


	So, it accepts a function of type HandlerFunc(req, res) and returns the same type

	so, here:
	rt.middlewares[i] is middleware of type Middleare and routeHandler is of type HandlerFunc

	So, Basically, it is:
		Middlware(HandlerFunc) {
			return HandlerFunc
		}

*/

import (
	"log"
	"net"
	"squirrel/core"
	"strings"
)

// every route will have a method:
// every route will have a pattern: /path/to
// every route will have a handler: func()
// every route may or may not have middlewares
type route struct {
	method      string
	pattern     string
	handler     core.HandlerFunc
	middlewares []Middleware // for route specific middlewares
}

type Middleware func(core.HandlerFunc) core.HandlerFunc

type SqurlMux struct {
	routes     []route
	middleware []Middleware // for global middlewares
}

func SpawnServer() *SqurlMux {
	return &SqurlMux{}
}

// method for global middleware handlers
// global middleware is handled by MuxServer
func (sm *SqurlMux) Use(mw Middleware) {
	sm.middleware = append(sm.middleware, mw)
}

// Now All the Methods need to be able to
// accept middlwares, after the actual handler
// the middleware will be executed in same  order they are kept in
// in which they are placed
func (sm *SqurlMux) Get(path string, handler core.HandlerFunc, mws ...Middleware) {
	sm.routes = append(sm.routes, route{
		method:      "GET",
		pattern:     path,
		handler:     handler,
		middlewares: mws, // route specific middlewares
	})

}

func (sm *SqurlMux) Post(path string, handler core.HandlerFunc, mws ...Middleware) {

	sm.routes = append(sm.routes, route{
		method:      "POST",
		pattern:     path,
		handler:     handler,
		middlewares: mws, // route specific middlewares
	})
}

func (sm *SqurlMux) Put(path string, handler core.HandlerFunc, mws ...Middleware) {

	sm.routes = append(sm.routes, route{
		method:      "PUT",
		pattern:     path,
		handler:     handler,
		middlewares: mws, // route specific middlewares
	})
}

func (sm *SqurlMux) Delete(path string, handler core.HandlerFunc, mws ...Middleware) {

	sm.routes = append(sm.routes, route{
		method:      "DELETE",
		pattern:     path,
		handler:     handler,
		middlewares: mws, // route specific middlewares
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

					// adding middleware to the respective handlers
					routeHandler := rt.handler

					// first handle route specific  middleware  handlers
					// cause we need to handle the global middleware first
					for i := len(rt.middlewares) - 1; i >= 0; i-- {
						// explanation at top
						routeHandler = rt.middlewares[i](routeHandler) // Logger(routeHandler)
					}

					// now handle global middlewares specific handlers
					// cause it needs to be executed first
					for i := len(sm.middleware) - 1; i >= 0; i-- {
						routeHandler = sm.middleware[i](routeHandler) // keep on wrapping the middleware
					}

					routeHandler(req, res)
					res.Send()
					return
				}
			}
			res.SetStatus(404)
			res.Write("Route not Found")
			res.Send()

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
