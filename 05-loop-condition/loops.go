package main

import (
	"fmt"
	"time"
)

func main1() {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("n:", n)

	maps := map[string]int{"Tehran": 1, "Karaj": 2, "Gorgan": 3, "Rasht": 4}
	maps["Isfahan"] = 5

	for key, value := range maps {
		fmt.Println("key:", key, "| value:", value)
	}

	c := 0
	for {
		c++
		if c%2 != 0 {
			continue
		}
		if c > 10 {
			break
		}
		fmt.Printf("%d Infinite loop ...\n", c)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("After Loop!")
}
