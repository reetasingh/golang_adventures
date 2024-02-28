package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	Name string
}

type T interface {
	A | B
}

func getName[t T](val t) string {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Struct {
		field := v.FieldByName("Name")
		if field.IsValid() {
			return field.String()
		}
	}
	return ""
}

func displayName[t T](val t) {
	fmt.Println(getName(val))
}

// displayName by directly accessing field
// this gives compile time error
// val.Name undefined (type t has no field or method Name)
// func displayName[t T](val t) {
// 	fmt.Println(val.Name)
// }

// displayName by type assertion
// this gives compile time error
// cannot use type switch on type parameter value val (variable of type t constrained by T)
// func displayName[t T](val t) {
// 	switch val.(type) {
// 	case A:
// 		fmt.Println(val.Name)
// 	case B:
// 		fmt.Println(val.Name)
// 	}
// }

func main() {
	a := A{"A struct"}
	b := B{"B struct"}
	//c := C{"C struct"}
	displayName(a)
	displayName(b)
	//displayName(c)
}
