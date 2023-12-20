package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var data []byte
var lock sync.RWMutex
var wg sync.WaitGroup

func produce() {
	lock.Lock()
	defer lock.Unlock()
	for i := 0; i < 10; i++ {
		data = append(data, byte(rand.Int()))
	}
	defer wg.Done()
}

func consume(i int) {
	lock.RLock()
	defer lock.RUnlock()
	if len(data) > 0 {
		fmt.Printf("\ngoroutine %d read %v, %v", i, data, time.Now())
	}
	defer wg.Done()
}

func main() {
	go produce()
	wg.Add(1)
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		go consume(i)
		wg.Add(1)
	}

	wg.Wait()

}
