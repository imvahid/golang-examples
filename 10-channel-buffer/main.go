package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
	for {
		v := <-pChannel
		fmt.Println("========================")
		fmt.Println(v)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	logChannel := make(chan string, 10)

	go printLog(logChannel)

	for i := 0; i < 20; i++ {
		fmt.Println("⬇ pushing in channel:", i)
		logChannel <- fmt.Sprintf("⬅ pulled form channel: %d", i)
		fmt.Println("➡ pushed  in channel:", i)
	}
	fmt.Println("✔ finished pushing in channel")

	time.Sleep(10 * time.Second)
}
