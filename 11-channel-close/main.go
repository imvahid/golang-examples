package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
	for {
		v, ok := <-pChannel
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println(v)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	logChannel := make(chan string, 10)

	go printLog(logChannel)

	for i := 0; i < 10; i++ {
		logChannel <- fmt.Sprintf("counter is: %d", i)
	}
	close(logChannel)

	time.Sleep(10 * time.Second)
}
