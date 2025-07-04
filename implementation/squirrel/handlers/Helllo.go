package handlers

import (
	"fmt"
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

	url := req.URL
	id := req.Params["id"]
	fmt.Println("-------")
	fmt.Println("url: ", url)
	fmt.Println("id: ", id)
	fmt.Println("-------")
	res.SetHeader("X-Use-B	ytes", "Using Bytes")
	res.WriteBytes([]byte("What the hell is happening here??"))
}

func BasicMiddleware(next core.HandlerFunc) core.HandlerFunc {
	return func(r1 *core.Request, r2 *core.Response) {
		fmt.Println("Woo, it working")
		next(r1, r2)
		fmt.Println("woo i passed")
	}
}

func CheckFile(next core.HandlerFunc) core.HandlerFunc {
	return func(rq *core.Request, rs *core.Response) {
		fmt.Println(rq.Path, rq.Method)
		next(rq, rs)
	}
}

func ReadFile(req *core.Request, res *core.Response) {
	file, err := os.Open("./core/response.go")
	if err != nil {
		log.Fatal("Error opening file")
	}
	res.SetHeader("Content-Type", "text/plain")
	res.SetBody(file)
}
