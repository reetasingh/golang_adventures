package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var data []byte
var lock sync.Mutex

func produce() {
	lock.Lock()
	defer lock.Unlock()
	time.Sleep(1 * time.Second)
	data = append(data, byte(rand.Int()))
}

func consume() {
	lock.Lock()
	defer lock.Unlock()
	if len(data) > 0 {
		elem := data[0]
		fmt.Println(elem)
		data = data[1:]
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go produce()
	}
	for i := 0; i < 10; i++ {
		go consume()
	}
	time.Sleep(1 * time.Minute)
}
