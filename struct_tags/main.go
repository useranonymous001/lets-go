package main

// to attach meta-information to the field which can be acquired using reflection.

// The key usually denotes the package that the subsequent "value"
// is for, for example json keys are processed/used by the encoding/json package.
type User struct {
	Name string `json:"name"`
}

func main() {
}
