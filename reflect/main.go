package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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

func ParseEnvFile(filePath string) ([]string, error) {
	var output []string

	if filePath == "" {
		return output, nil
	}

	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		return output, err
	}
	_, err = os.Stat(fullPath)
	if err != nil {
		return output, err
	}

	fileData, err := os.ReadFile(fullPath)
	if err != nil {
		return output, err
	}

	if len(fileData) == 0 {
		return output, nil
	}

	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				output = append(output, line)
			}
		}
	}
	return output, nil
}

func main() {
	//example := Example{10, 21.1, []string{"demo", "value"}, 20}
	//analyze(example)
	result, err := ParseEnvFile("reflect/abc.env")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
