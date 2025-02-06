package main

import (
	"fmt"
	"sync"
)

/* Go FORK-JOIN Model */

/*
the anonymous function will be forked but there is no join point
So, it is undetermined if sayHello function will be executed before main goroutine exits
*/

/*
can be solved using sync package to create join point
*/
func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Concurrency in Go")
	}
	wg.Add(1)
	go sayHello()
	fmt.Println("Hello World")
	wg.Wait()
}
