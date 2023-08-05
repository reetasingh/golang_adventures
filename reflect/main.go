package main

import (
	"fmt"
	"reflect"
)

type Myint int
type Example struct {
	a int64
	b float32
	c []string
	d Myint
}

func analyze[T any](elem T) {
	fmt.Println(reflect.TypeOf(elem)) // Example
	fmt.Println(reflect.ValueOf(elem))
	fmt.Println(reflect.ValueOf(elem).String())

	fmt.Println(reflect.ValueOf(elem).Kind()) // struct, kind refers to underlying type

	// read the fields and their type and value
	s := reflect.ValueOf(&elem).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println("-----")
		fmt.Println(s.Field(i).Type())
		fmt.Println(s.Field(i).Type().Name())
		fmt.Println(s.Field(i).Kind())
		fmt.Println(s.Field(i))
	}
}

func main() {
	example := Example{10, 21.1, []string{"demo", "value"}, 20}
	analyze(example)
}
