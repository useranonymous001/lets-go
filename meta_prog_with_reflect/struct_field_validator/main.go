// Project 2: Struct Field Validator using Tags
// custom validator to validate the struct tags using reflect

/*
Currently, the custom validator supports for the following validations:

"tag: required": the field should not be empty
"tag: min=": min int value to enter
"tag: max=": max int value that will be supported

*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string `validate:"required"`
	Age  int    `validate:"min=16"`
	// Age      int    `validate:"min=18 max=10"`
	Salary int `validate:"max=1000"`
}

func main() {

	user := User{
		Name:   "rohan",
		Age:    17,
		Salary: 100,
	}

	errs := Validator(&user)

	for _, err := range errs {
		fmt.Println(err)
	}
}

func Validator(data interface{}) []error {

	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	// collect errors
	var errs []error

	// parse the struct and get the tag out of it.
	for i := 0; i < typ.NumField(); i++ {

		field := typ.Field(i)
		value := val.Field(i) // represents the value of fields in struct

		tag := field.Tag.Get("validate") // required || "" || min= || max=

		// check for the struct Tags implemented or not
		if tag == "" {
			// leave this Tag and continue
			continue
		}

		// if so, check for the value

		if tag == "required" {
			if isZero(value) {
				errs = append(errs, fmt.Errorf("Field %s is required", typ.Field(i).Name))
			}
		}

		// check the tag has --min-- field
		if strings.HasPrefix(tag, "min=") {
			// cut out the min=
			minValStr := strings.TrimPrefix(tag, "min=")

			// convert the rest to int
			minValInt, _ := strconv.Atoi(minValStr)

			if value.Kind() == reflect.Int && value.Int() < int64(minValInt) {
				errs = append(errs, fmt.Errorf("Field %s must be at least %d", field.Name, minValInt))
			}
		}

		// check the tag has --max-- field
		if strings.HasPrefix(tag, "max=") {
			// get the max= prefix
			maxStrVal := strings.TrimPrefix(tag, "max=")

			// convert it into integers now:
			maxIntVal, _ := strconv.Atoi(maxStrVal)

			if value.Kind() == reflect.Int && value.Int() > int64(maxIntVal) {
				errs = append(errs, fmt.Errorf("Field %s must not be more than %d", field.Name, maxIntVal))
			}

		}

	}
	return errs
}

func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

	