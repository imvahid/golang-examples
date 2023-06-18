package main

import "fmt"

type person struct {
	name   string
	family string
	age    int
}

func main() {
	var p person

	p.name = "Vahid"
	p.family = "Ashourzadeh"
	p.age = 32

	fmt.Println("p:", p)

	d := person{}
	fmt.Println("d:", d)

	f := person{family: "foo", age: 23}
	fmt.Println("f:", f)
}
