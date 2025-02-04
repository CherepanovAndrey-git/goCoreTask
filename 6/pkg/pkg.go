package pkg

import (
	"math/rand"
	"time"
)

func RandomNumGenerator(nums chan<- int, done <-chan struct{}) {
	defer close(nums)

	for {
		select {
		case <-done:
			return
		default:
			select {
			case nums <- rand.Intn(1000):
				time.Sleep(500 * time.Millisecond)
			case <-done:
				return
			}
		}
	}
}
