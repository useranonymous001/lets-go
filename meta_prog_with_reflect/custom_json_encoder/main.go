/*
*** Custom JSON encoder using reflect package.
## convert struct fields into json

*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name      string
	Email     string
	Password  string
	Age       int
	Roles     []string
	Is_active bool
	Salary    float64
}

// parse the struct and then convert each field into json field
func main() {

	// let's create a instance of an User struct
	u := User{
		Name:      "gwen stacy",
		Email:     "gwen@stacy.com",
		Password:  "gwenthegreatgirl",
		Age:       19,
		Roles:     []string{"sr. software engineer", "hr"},
		Is_active: true,
		Salary:    891000.12,
	}

	// Test(u)
	// get the kind of each field type
	EncodeToJson(u)
	// fmt.Println(jsonStr)

	Conversion()

}

// encoder func that encodes the struct to json
func EncodeToJson(data interface{}) string {

	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)

	// check the type is pointer or not
	if val.Kind() == reflect.Ptr {
		typ = typ.Elem() // dereference to get the underlying value
		val = val.Elem() // same thing here..
	}

	// check whether the reflected type is struct or not
	if val.Kind() != reflect.Struct {
		return `{}` // empty json
	}

	jsonStr := "{"

	for i := 0; i < val.NumField(); i++ {

		// get the field and it's value
		field := typ.Field(i)
		value := val.Field(i)

		// value.Interface()       // value of that field
		// log.Println(field.Name) // name of that field

		// concatenate strings checking thier each type
		jsonStr += `"` + field.Name + `":`

		switch value.Kind() {
		case reflect.String:
			jsonStr += `"` + value.String() + `"`

		case reflect.Int, reflect.Int32, reflect.Int8, reflect.Int16, reflect.Int64:
			jsonStr += strconv.FormatInt(value.Int(), 10)

		case reflect.Bool:
			jsonStr += strconv.FormatBool(value.Bool())

		case reflect.Float32, reflect.Float64:
			jsonStr += strconv.FormatFloat(value.Float(), 'f', -1, 64)

		default:
			jsonStr += `"unsupported"`

		}

		// add comma if it's not the last field
		if i < val.NumField()-1 {
			jsonStr += ","
		}

	}
	jsonStr += "}"

	// after convers

	return jsonStr

}

func Conversion() {

	// var value int32 = 123

	// fmt.Println(strconv.FormatInt(int64(value), 10))
	// fmt.Printf("%T", strconv.FormatInt(int64(value), 10))

	// var isActive bool = true
	// fmt.Println(strconv.FormatBool(isActive))
	// fmt.Printf("%T", isActive)

	var value int64 = 123
	// strVal := strconv.FormatInt(value, 10)
	// fmt.Printf("%T\n", strVal)

	typ := reflect.TypeOf(&value)
	val := reflect.ValueOf(&value)

	typ = typ.Elem()
	val = val.Elem()

	strVal := val.String()

	val.SetInt(1324)
	fmt.Println(strVal)
	fmt.Println(val.Interface())
	fmt.Println("actual value:", value)
}

func Test(data interface{}) string {
	// we reflect the type of user struct
	typ := reflect.TypeOf(data)  // reflects the type of variable
	val := reflect.ValueOf(data) // reflects the underlying value of the type

	// type includes
	fmt.Println(typ.Name())                               // name of the type
	fmt.Println(typ.Kind())                               // kind of the type
	fmt.Println(typ.NumField())                           // number of fields inside the structure
	fmt.Println("typ.Field(0).Name: ", typ.Field(0).Name) // name of fields in that type

	// value includes
	fmt.Println("val.Field(0): ", val.Field(0)) // prints the reflect.Value of the Name
	fmt.Println(val.Interface())                // prints the actual value of the Name

	for i := 0; i < val.NumField(); i++ {
		fmt.Println(val.Field(i).Interface())
	}

	return ""
}
