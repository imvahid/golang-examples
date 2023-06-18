package main

import "fmt"

func main() {
	i := 52

	var p *int
	fmt.Println("p:", p)
	p = &i

	fmt.Println("i:", i)
	fmt.Println("p:", p)
	fmt.Println("&i:", &i)
	fmt.Println("*p:", *p)

	i = 92

	fmt.Println("*p:", *p, "p:", p)
}
