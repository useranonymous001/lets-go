package main

import (
	"squirrel/handlers"
	"squirrel/server"
)

func main() {

	server := server.SpawnServer()
	server.Post("/hello", handlers.Hello)
	server.Listen(":9000")
}
