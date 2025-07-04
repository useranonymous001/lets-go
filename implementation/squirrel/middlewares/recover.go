package middlewares

import (
	"fmt"
	"log"
	"squirrel/core"
)

func Recover(next core.HandlerFunc) core.HandlerFunc {
	return func(req *core.Request, res *core.Response) {
		str := recover()
		if str != nil {
			func() {
				log.Fatal(str)
			}()
		}
		next(req, res)
	}
}

func Logger(next core.HandlerFunc) core.HandlerFunc {
	return func(req *core.Request, res *core.Response) {
		fmt.Println(req.Path, req.Method)
		next(req, res)
	}
}
