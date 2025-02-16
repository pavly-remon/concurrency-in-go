package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("------------------Example1----------------\n")
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.", time.Since(start))
	}
	fmt.Printf("\n\n------------------Example2----------------\n")
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count:%d\nc2Count:%d\n", c1Count, c2Count)
}
