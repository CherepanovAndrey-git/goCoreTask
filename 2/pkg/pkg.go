package pkg

// SliceExample Принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
func SliceExample(slice []int) []int {
	newSlice := []int{}

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 && slice[i] != 0 {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

// AddElements Принимает слайс и число. Функция добавляет это число в конец слайса и возвращает новый слайс.
func AddElements(slice []int, num int) []int {
	newSlice := make([]int, len(slice)+1)

	for i, v := range slice {
		newSlice[i] = v
	}

	newSlice[len(slice)] = num

	return newSlice
}

// CopySlice принимает слайс, возвращает копию.
func CopySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	for i, v := range slice {
		newSlice[i] = v
	}
	return newSlice
}

// RemoveElement принимает слайс и индекс элемента который требуется удалить, возвращает новый слайс, без данных в указаном индексе
func RemoveElement(s []int, index int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, s[:index]...)
	return append(newSlice, s[index+1:]...)
}
