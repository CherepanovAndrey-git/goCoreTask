package main

import (
	"9/pkg"
	"fmt"
)

func main() {
	inputChan := make(chan uint8)
	outputChan := pkg.CubeConv(inputChan)

	go func() {
		defer close(inputChan)
		numbers := []uint8{2, 3, 4, 5, 10}
		for _, num := range numbers {
			inputChan <- num
		}
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
