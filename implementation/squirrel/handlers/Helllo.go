package handlers

import (
	"fmt"
	"squirrel/core"
)

func Hello(req *core.Request, res *core.Response) {
	res.SetHeader("X-Squirrel", "Lazi Squirrel Http Framework")

	fmt.Println(req.Body)
	fmt.Println(req.URL)
	fmt.Println(req.Headers)
	fmt.Println(req.Path)
	fmt.Println(req.Method)
	fmt.Println(req.ContentLength)

	res.Write("Hello From Squirrle HTTP Framework")
}
