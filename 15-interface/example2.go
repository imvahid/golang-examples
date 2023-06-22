package main

import "fmt"

type Shape interface {
	Area() int
}

type Square struct {
	Length int
}

func (s Square) Area() int {
	return s.Length * s.Length
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func example2() {
	s := Square{Length: 10}
	r := Rectangle{Length: 10, Width: 20}

	shapeList := []Shape{s, r}

	for _, value := range shapeList {
		fmt.Printf("type: %T, area: %d\n", value, value.Area())

		// type assertion
		square, ok := value.(Square)
		fmt.Printf("type: %T, ok: %v\n", square, ok)
		fmt.Println("..................")
	}

	sArea := CalculateArea(s)
	fmt.Println("CalculateArea(s):", sArea)
}

func CalculateArea(s Shape) int {
	return s.Area()
}
