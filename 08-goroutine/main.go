package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go say("Hello", &wg)
	wg.Wait()

	wg.Add(1)
	go say("World", &wg)
	wg.Wait()
}
