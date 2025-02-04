package pkg

func CubeConv(input <-chan uint8) <-chan float64 {
	output := make(chan float64)
	go func() {
		defer close(output)
		for num := range input {
			f := float64(num)
			output <- f * f * f
		}
	}()
	return output
}
