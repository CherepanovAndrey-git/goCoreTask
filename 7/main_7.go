package main

import (
	"7/pkg"
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 11; i <= 13; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 21; i <= 23; i++ {
			ch3 <- i
		}
	}()

	merged := pkg.MergeChannels(ch1, ch2, ch3)

	for num := range merged {
		fmt.Println(num)
	}
}
