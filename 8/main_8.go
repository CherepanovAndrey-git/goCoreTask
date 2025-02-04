package main

import (
	"sync"
	"sync/atomic"
)

// CustomWaitGroup аналог sync.WaitGroup с семафором
type CustomWaitGroup struct {
	counter atomic.Int64
	sem     chan struct{}
	mu      sync.Mutex
}

// NewCustomWaitGroup создает новый экземпляр CustomWaitGroup.
func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem: make(chan struct{}),
	}
}

// Add добавляет дельту к счетчику.
func (wg *CustomWaitGroup) Add(delta int) {
	newCounter := wg.counter.Add(int64(delta))
	if newCounter < 0 {
		panic("negative counter")
	}

	wg.mu.Lock()
	defer wg.mu.Unlock()

	if newCounter == 0 {
		close(wg.sem)
	} else if delta > 0 && newCounter == int64(delta) {
		wg.sem = make(chan struct{})
	}
}

// Done уменьшает счетчик на 1.
func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокируется, пока счетчик не станет нулевым.
func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	sem := wg.sem
	wg.mu.Unlock()

	<-sem
}
