package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestCustomWaitGroupSingle(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter = 42
	}()

	wg.Wait()

	if counter != 42 {
		t.Errorf("Expected 42, got %d", counter)
	}
}

func TestCustomWaitGroupMultiple(t *testing.T) {
	wg := NewCustomWaitGroup()
	const n = 5
	var counter atomic.Int32

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	if cnt := counter.Load(); cnt != n {
		t.Errorf("Expected %d, got %d", n, cnt)
	}
}

func TestCustomWaitGroupReuse(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(2)
	go func() { wg.Done() }()
	go func() { wg.Done() }()
	wg.Wait()

	var counter atomic.Int32
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()

	if cnt := counter.Load(); cnt != 3 {
		t.Errorf("Expected 3, got %d", cnt)
	}
}

func TestNegativeCounter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative counter")
		}
	}()

	wg := NewCustomWaitGroup()
	wg.Add(-1)
}

func TestConcurrentWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	const n = 100
	wg.Add(n)

	done := make(chan struct{})

	go func() {
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Timeout waiting for Wait()")
	}
}
