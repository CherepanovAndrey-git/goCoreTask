package main

import (
	"6/pkg"
	"fmt"
)

func main() {
	numsChan := make(chan int)
	done := make(chan struct{})

	go pkg.RandomNumGenerator(numsChan, done)

	for i := 0; i < 100; i++ {
		number := <-numsChan
		fmt.Printf("Got number: %d\n", number)
	}

	close(done)
	fmt.Println("Random number generator has stopped.")

}
