package main

import (
	"squirrel/handlers"
	"squirrel/server"
)

func main() {

	server := server.SpawnServer()
	server.Get("/hello", handlers.Hello)
	server.Listen(":9000")
}
