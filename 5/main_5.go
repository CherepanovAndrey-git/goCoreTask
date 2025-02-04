package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{65, 3, 58, 678, 64, 12}
	b := []int{64, 2, 3, 43, 12}
	fmt.Println(FindIntersection(a, b))
}

// FindIntersection Функция проверяет пересечения и возвращает результат, используем карту
func FindIntersection(a, b []int) (bool, []int) {
	elements := make(map[int]bool)
	for _, num := range b {
		elements[num] = true
	}

	added := make(map[int]bool)
	intersection := []int{}

	for _, num := range a {
		if elements[num] && !added[num] {
			intersection = append(intersection, num)
			added[num] = true
		}
	}
	sort.Ints(intersection)

	return len(intersection) > 0, intersection
}
