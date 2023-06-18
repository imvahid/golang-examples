package main

import "fmt"

func main() {
	checkMultiple(30)

	dayOfWeek(1)
}

func checkMultiple(x int) {
	s := []int{}

	if x%2 == 0 {
		s = append(s, 2)
	}
	if x%3 == 0 {
		s = append(s, 3)
	}
	if x%4 == 0 {
		s = append(s, 4)
	}
	if x%5 == 0 {
		s = append(s, 5)
	}
	if x%6 == 0 {
		s = append(s, 6)
	}
	if x%7 == 0 {
		s = append(s, 7)
	}
	if x%8 == 0 {
		s = append(s, 8)
	}
	if x%9 == 0 {
		s = append(s, 9)
	}

	fmt.Printf("%d is multiple of: %v\n", x, s)
}

func dayOfWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day!")
	}
}
