package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStack(name string) {
	buf := make([]byte, 1024)
	stackSize := runtime.Stack(buf, false)
	fmt.Printf("%s %s\n", name, buf[:stackSize])
}

func main() {
	go func() {
		// Some goroutine work
		time.Sleep(1 * time.Second)
		printStack("abc")
	}()

	go func() {
		// Some goroutine work
		time.Sleep(1 * time.Second)
		printStack("pqr")
	}()

	// Some main goroutine work
	time.Sleep(3 * time.Second)
}
