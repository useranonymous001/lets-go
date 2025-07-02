package main // tells the compiler to run this program as a standalone program

import (
	"fmt"
	"strings"
)

// this is comment
func main() {

	pattern := "/byte/:id/hello"
	path := "/byte/123/hello"
	params := map[string]string{}

	fmt.Println(matchPattern(pattern, path, params))

}

func matchPattern(pattern, path string, params map[string]string) bool {
	patternSegment := strings.Split(pattern, "/")
	pathSegment := strings.Split(path, "/")

	if len(pathSegment) != len(patternSegment) {
		return false
	}

	for i := 0; i < len(patternSegment); i++ {
		fmt.Println(patternSegment[i])
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
