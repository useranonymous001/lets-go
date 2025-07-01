package handlers

import (
	"log"
	"os"
	"squirrel/core"
)

func Hello(req *core.Request, res *core.Response) {
	res.SetHeader("X-Squirrel", "Lazi Squirrel Http Framework")
	data, _ := req.ReadBodyAsString()
	res.Write(data)
}

func UseJson(req *core.Request, res *core.Response) {
	res.SetHeader("X-Use-Json", "Demo")
	res.JSON(map[string]string{
		"message":  "hello",
		"notebook": "writing",
	})
}

func UseBytes(req *core.Request, res *core.Response) {
	res.SetHeader("X-Use-Bytes", "Using Bytes")
	res.WriteBytes([]byte("What the hell is happening here??"))
}

func ReadFile(req *core.Request, res *core.Response) {
	file, err := os.Open("./core/response.go")
	if err != nil {
		log.Fatal("Error opening file")
	}
	res.SetHeader("Content-Type", "text/plain")
	res.SetBody(file)
}
