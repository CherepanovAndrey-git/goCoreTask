package main

import (
	"6/pkg"
	"context"
	"sync"
	"testing"
	"time"
)

func TestRandomNumGenerator_BasicFunctionality(t *testing.T) {
	numsChan := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		pkg.RandomNumGenerator(numsChan, done)
	}()

	for i := 0; i < 5; i++ {
		select {
		case num, ok := <-numsChan:
			if !ok {
				t.Fatal("Channel closed prematurely")
			}
			if num < 0 || num >= 1000 {
				t.Errorf("Number %d out of range [0, 1000)", num)
			}
		case <-time.After(2 * time.Second):
			t.Fatal("Timeout waiting for number")
		}
	}

	close(done)
	wg.Wait()

	for {
		select {
		case num, ok := <-numsChan:
			if !ok {
				return
			}
			if num < 0 || num >= 1000 {
				t.Errorf("Number %d out of range [0, 1000)", num)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Channel should be closed")
			return
		}
	}
}

func TestRandomNumGenerator_ImmediateStop(t *testing.T) {
	numsChan := make(chan int)
	done := make(chan struct{})
	close(done)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		pkg.RandomNumGenerator(numsChan, done)
	}()

	wg.Wait()

	select {
	case _, ok := <-numsChan:
		if ok {
			t.Error("Channel should be closed")
		}
	default:
		t.Error("Channel is not closed")
	}
}

func TestRandomNumGenerator_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	numsChan := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		pkg.RandomNumGenerator(numsChan, done)
	}()

Loop:
	for {
		select {
		case num, ok := <-numsChan:
			if !ok {
				break Loop
			}
			if num < 0 || num >= 1000 {
				t.Errorf("Number %d out of range", num)
			}
		case <-ctx.Done():
			close(done)
			break Loop
		}
	}

	wg.Wait()
}
