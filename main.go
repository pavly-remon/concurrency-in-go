package main

import "fmt"

func chanOwner() <-chan int {
	c := make(chan int, 3)
	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	return c
}
func main() {
	outputStream := chanOwner()
	for v := range outputStream {
		fmt.Printf("Received: %d\n", v)
	}
	fmt.Println("Done Receiving")
}
