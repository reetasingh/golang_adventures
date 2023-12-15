package main

import (
	"errors"
	"fmt"
)

var NotFoundErr = errors.New("key not found")

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
	// check error type using interface type check
	if err != nil {
		err = &MyError{"test"}
		if myErr, ok := err.(*MyError); ok {
			fmt.Printf("Is Unwrapped Error: %v\n", myErr.Message)
		} else {
			fmt.Printf("Not is Unwrapped Error: %v\n", err)
		}
	}

	// unwrap errors till the end
	lastErr := getLastError(err)
	fmt.Printf("\nunwrapped last error %v\n", lastErr)

	// example of using errors.Is for checking error
	err = dbErrorFunction()
	if err != nil {
		if errors.Is(err, NotFoundErr) {
			fmt.Printf("error matches %v", err)
		} else {
			fmt.Printf("error does not matches")
		}
	}
	// unwrap errors till the end
	lastErr = getLastError(err)
	fmt.Printf("\nunwrapped last error %v", lastErr)

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

func dbErrorFunction() error {
	//return fmt.Errorf("random error")
	//return randomErr
	return fmt.Errorf("error chain %w", NotFoundErr)
}

// unwrap errors till the last error
func getLastError(err error) error {
	newErr := errors.Unwrap(err)
	if newErr == nil {
		return err
	} else {
		return getLastError(newErr)
	}
}
