package main

import (
	"fmt"
	"sync"
)

type Container struct {
	counters map[string]int // shared resource that doesn't support concurrent writes
	mu       sync.Mutex     // protect the shared resource with a lock
}

func (c *Container) inc(name string) {
	// try removing the lock here and see what happens
	// the program will panic with an error
	// fatal error: concurrent map writes
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Container) dec(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]--
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0}, // map doesn't support concurrent access
	}

	var wg sync.WaitGroup

	increment := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name) // increment counter
		}
		wg.Done()
	}

	decrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.dec(name) // decrement counter
		}
		wg.Done()
	}

	// if you WaitGroup counter is lower than needed your program will panic
	// if it's higher it will be blocked in a deadlock
	wg.Add(3)
	go decrement("a", 1000)
	go increment("a", 1000)
	go increment("b", 1000)

	wg.Wait()

	fmt.Println("Final Counters:", c.counters)
}
