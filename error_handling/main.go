package main

import (
	"errors"
	"fmt"
)

var randomErr = errors.New("random error")

func main() {
	// this debug msg will now be logged
	err := someFunction()
	// example of using errors.As for checking error type
	if err != nil {
		var unwrappedErr *MyError
		if errors.As(err, &unwrappedErr) {
			fmt.Printf("Unwrapped Error: %v\n", unwrappedErr.Message)
		} else {
			fmt.Println("Error is not of type *MyError")
		}
	}

	// example of using errors.Is for checking error
	err = randomErrorFunction()
	if err != nil {
		if errors.Is(err, randomErr) {
			fmt.Printf("error matches %v", err)
		} else {
			fmt.Printf("error does not matches")
		}
	}
}

type MyError struct {
	Message string
}

func (myError *MyError) Error() string {
	return myError.Message
}

func someFunction() error {
	// Simulating an error
	err := someOtherFunction()
	if err != nil {
		return fmt.Errorf("db error enountered %w", err)
	}
	return nil
}

func someOtherFunction() error {
	// Simulating an error
	return &MyError{Message: "this is db error"}
}

func randomErrorFunction() error {
	//return fmt.Errorf("random error")
	//return randomErr
	return fmt.Errorf("random error chain %w", randomErr)
}
