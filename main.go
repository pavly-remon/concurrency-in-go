package main

import "fmt"

/*
According to fork-join model:
the anonymous function will be forked but there is no join point
So, it is undetermined if sayHello function will be executed before main goroutine exits
*/
func main() {
	sayHello := func() {
		fmt.Println("Concurrency in Go")
	}
	go sayHello()
	fmt.Println("Hello World")
}
