package main

import (
	"fmt"
	"time"
)

func printer(ch chan interface{}) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Printf("channel closed")
			return
		}

		switch t := v.(type) {
		case int:
			fmt.Printf("type is `%T`, value: %d\n", t, v)
		case string:
			fmt.Printf("type is `%T`, value: %s\n", t, v)
		case float64:
			fmt.Printf("type is `%T`, value: %v\n", t, v)
		case map[string]int:
			fmt.Printf("type is `%T`, value: %v\n", t, v)
		case []int:
			fmt.Printf("type is `%T`, value: %v\n", t, v)
		case Square:
			fmt.Printf("type is `Square`, value: Length => %d\n", v.(Square).Length)
		case Rectangle:
			fmt.Printf("type is `Rectangle`, value: Length => %d, Width => %d\n", v.(Rectangle).Length, v.(Rectangle).Width)
		default:
			fmt.Printf("type `%T` not supported, value: %v\n", t, v)
		}
	}
}

func example3() {
	ch := make(chan interface{}, 10)

	s := Square{Length: 10}
	r := Rectangle{Length: 20, Width: 30}

	go printer(ch)

	ch <- "string"
	ch <- 123455
	ch <- 1.2345
	ch <- []int{1, 2}
	ch <- map[string]int{"key1": 1, "key2": 2}
	ch <- s
	ch <- r

	time.Sleep(1 * time.Second)
}
