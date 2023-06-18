package printer

import (
	"fmt"
	"time"
)

// PrintHelloWorld method is print `Hello World` message with time
func PrintHelloWorld() {
	printTime()
	fmt.Println("Hello World!")
}

func printTime() {
	fmt.Print(time.Now().Format("2006-04-02 15:04:05"))
	fmt.Print(" => ")
}
