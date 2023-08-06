package main

import (
	"flag"
	"fmt"
	"reflect"
)

type MyStruct struct {
}

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
)

func processTag(tag reflect.StructTag) {
	val := tag.Get("json")
	fmt.Println(val)
}

func convertTag(tag reflect.StructTag) reflect.StructTag {
	val := tag.Get("json")
	s := fmt.Sprintf(`db:"%s"`, val)
	newDBTag := reflect.StructTag(s)
	return newDBTag
}

// CreateDynamicStruct creates a new struct type dynamically based on the fields of the source type
func CreateDynamicStruct(sourceType reflect.Type, structName string) reflect.Type {
	numFields := sourceType.NumField()
	fieldTypes := make([]reflect.StructField, numFields)

	for i := 0; i < numFields; i++ {
		field := sourceType.Field(i)
		if field.Type.Kind().String() == "struct" {
			fieldTypes[i] = reflect.StructField{
				Name: field.Name,
				Type: CreateDynamicStruct(field.Type, field.Name),
			}
		} else {
			fieldTypes[i] = reflect.StructField{
				Name: field.Name,
				Type: field.Type,
				Tag:  convertTag(field.Tag),
			}
		}
	}

	newStructType := reflect.StructOf(fieldTypes)
	return newStructType
}

type ChildStruct struct {
	Field3 int `json:"field3,omitempty"`
}

// SourceStruct represents the source struct
type SourceStruct struct {
	Field1 int    `json:"field1,omitempty"`
	Field2 string `json:"field2,omitempty"`
	Child  ChildStruct
	// Add other fields as needed
}

func main() {
	// Get the reflect.Type of SourceStruct
	sourceType := reflect.TypeOf(SourceStruct{})

	// Dynamically create a new struct type based on the fields of SourceStruct
	newStructType := CreateDynamicStruct(sourceType, "DynamicStruct")

	// Create an instance of the dynamically created struct type
	newStructValue := reflect.New(newStructType).Elem()

	// Print the type of the dynamically created struct
	fmt.Printf("Type: %T\nValue: %+v\n", newStructValue.Interface(), newStructValue.Interface())
}
