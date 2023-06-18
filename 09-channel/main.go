package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello ")
		pingChannel <- "World!"
		<-pongChannel
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func sayWorld(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	for i := 0; i < 5; i++ {
		v := <-pingChannel
		fmt.Printf("%s\n", v)
		pongChannel <- "Done"
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	pingChannel := make(chan string)
	pongChannel := make(chan string)

	wg.Add(1)
	go sayHello(&wg, pingChannel, pongChannel)

	wg.Add(1)
	go sayWorld(&wg, pingChannel, pongChannel)

	wg.Wait()

	fmt.Println("Done.")
}
