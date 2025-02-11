package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		defer cond.L.Unlock()
		defer cond.Signal()
		queue = queue[1:]
		fmt.Println("Removed from queue")
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for len(queue) == 2 {
			cond.Wait()
		}
		fmt.Println("Adding to Queue...")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		cond.L.Unlock()
	}
}
