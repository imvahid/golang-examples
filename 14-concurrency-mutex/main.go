package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at time can access the map c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at time can access the map c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	go func() {
		for i := 0; i < 15; i++ {
			c.Inc("https://example.com/login")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("https://example.com/news")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("https://example.com/login")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("https://example.com/login"))
}
