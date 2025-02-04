package pkg

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	t.Run("multiple channels", func(t *testing.T) {
		ch1 := generateChan(1, 3)
		ch2 := generateChan(4, 6)
		ch3 := generateChan(7, 9)

		merged := MergeChannels(ch1, ch2, ch3)

		var results []int
		for n := range merged {
			results = append(results, n)
		}

		if len(results) != 9 {
			t.Errorf("Expected 9 elements, got %d", len(results))
		}

		expected := map[int]bool{
			1: true, 2: true, 3: true,
			4: true, 5: true, 6: true,
			7: true, 8: true, 9: true,
		}

		for _, n := range results {
			if !expected[n] {
				t.Errorf("Unexpected value %d", n)
			}
		}
	})

	t.Run("single channel", func(t *testing.T) {
		ch := generateChan(10, 12)
		merged := MergeChannels(ch)

		var results []int
		for n := range merged {
			results = append(results, n)
		}

		expected := []int{10, 11, 12}
		if !slicesEqual(results, expected) {
			t.Errorf("Expected %v, got %v", expected, results)
		}
	})

	t.Run("no channels", func(t *testing.T) {
		merged := MergeChannels()
		select {
		case _, ok := <-merged:
			if ok {
				t.Error("Channel should be closed immediately")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected closed channel")
		}
	})

	t.Run("channel closure", func(t *testing.T) {
		ch := make(chan int)
		go func() {
			ch <- 1
			close(ch)
		}()

		merged := MergeChannels(ch)
		<-merged

		if _, ok := <-merged; ok {
			t.Error("Expected closed channel")
		}
	})
}

func generateChan(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
