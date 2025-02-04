package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Items in slice 1 that slice 2 doesn't have", SliceDiff(slice1, slice2))

}

func SliceDiff(slice1, slice2 []string) []string {
	//используем карту для проверки наличия элементов во втором слайсе, аппендим второй слайс только теми элементами которых нет
	m := make(map[string]bool)
	for _, s := range slice2 {
		m[s] = true
	}

	var diff []string
	for _, s := range slice1 {
		if !m[s] {
			diff = append(diff, s)
		}
	}
	return diff
}
