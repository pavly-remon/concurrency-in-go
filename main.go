package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	var lock sync.Mutex
	increment := func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Increment Count: %d\n", count)

	}
	decrement := func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrement Count: %d\n", count)

	}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go increment()
		go decrement()
	}
	wg.Wait()
}
