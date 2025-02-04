package main

import (
	"2/pkg"
	"math/rand"
)
import "fmt"

func main() {
	slice := rand.Perm(10) //генерируем срез из 10 случайных чисел
	fmt.Printf("original slice: %v\n", slice)

	fmt.Printf("even numbers slice from OG slice: %v\n", pkg.SliceExample(slice))
	copiedSlice := pkg.CopySlice(slice)         //копируем слайс
	newSlice := pkg.AddElements(copiedSlice, 2) // добавляем новый элемент в слайс, возвращается копия слайса
	copiedSlice[1] = 1234567890                 // меняем элемент в индексе, проверяем что оригинал не изменился
	fmt.Printf("copied slice: %v\n", copiedSlice)
	fmt.Printf("new slice after appending new element: %v\n", newSlice)
	fmt.Printf("_____________________________________\n")
	removedElSlice := pkg.RemoveElement(slice, 2)
	fmt.Printf("removed el slice: %v\n", removedElSlice)
	fmt.Printf("original slice: %v\n", slice)

}
