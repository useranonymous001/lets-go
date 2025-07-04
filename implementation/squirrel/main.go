package main

import (
	"squirrel/handlers"
	"squirrel/server"
)

func main() {

	server := server.SpawnServer()
	server.Post("/hello", handlers.Hello)
	server.Get("/file", handlers.ReadFile, handlers.CheckFile, handlers.BasicMiddleware)
	server.Get("/json", handlers.UseJson)
	server.Get("/byte/:id/read", handlers.UseBytes)
	server.Listen(":9000")

}
