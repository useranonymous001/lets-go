package handlers

import "squirrel/core"

func Hello(req *core.Request, res *core.Response) {
	res.SetHeader("X-Squirrel", "Lazi Squirrel Http Framework")
	res.Write("Hello From Squirrle HTTP Framework")
}
