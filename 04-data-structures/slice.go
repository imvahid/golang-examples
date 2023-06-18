package main

import "fmt"

var integerSlice []int
var stringSlice []string

func main3() {
	integerSlice = []int{10, 20, 30, 40}
	fmt.Println("integerSlice:", integerSlice)

	stringSlice = []string{"first", "second", "third"}
	fmt.Println("stringSlice:", stringSlice)

	s := []int{10, 20, 30, 40}
	fmt.Println("len(s):", len(s), "| cap(s):", cap(s))
	fmt.Println("s[0:2]:", s[0:2])
	fmt.Println("s[0:len(s)]:", s[0:len(s)])
	fmt.Println("s[0:]:", s[0:])
	fmt.Println("s[:2]:", s[:2])
	fmt.Println("s[:]:", s[:])

	for key, value := range s {
		fmt.Println("key:", key, "value:", value)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("i:", i, "s[", i, "]:", s[i])
	}

	// add
	s = append(s, 50)
	s = append(s, 60, 70)
	fmt.Println("s:", s)
}
