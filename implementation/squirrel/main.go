package main

import (
	"squirrel/handlers"
	"squirrel/server"
)

func main() {

	server := server.SpawnServer()
	server.Post("/hello", handlers.Hello)
	server.Get("/file", handlers.ReadFile)
	server.Get("/json", handlers.UseJson)
	server.Get("/byte", handlers.UseBytes)
	server.Listen(":9000")
}
