package main

import (
	"fmt"
	"math"
)

const float64EqualityThreshold = 1e-9

const (
	water = iota
	air
	soil
	glass
)

type float float64

func main() {
	var integer1 int
	fmt.Println(integer1)

	var string1 string
	string1 = "This is a string variable"
	fmt.Println(string1)

	var float1 = -1.2
	fmt.Println(float1)

	integer2 := 52
	fmt.Println(integer2)

	string2 := "This is a string"
	fmt.Println(string2)

	var var1, var2, var3 = -1.2, 1, false
	fmt.Println(var1, var2, var3)

	var (
		var4 float64 = -1.2
		var5 int     = 1
		var6 bool    = true
		var7 string
	)
	fmt.Println(var4, var5, var6, var7)

	// type conversion
	var8 := var4 + float64(var5)
	fmt.Println(var8)

	// IEEE 754
	if math.Abs(var8-(-0.2)) < float64EqualityThreshold {
		fmt.Println("It's Equal!")
	}

	//var f float = 5.2
	//var g float64 = 5.2
	//fmt.Println("f == g", f == g)

	fmt.Println(water, air, soil, glass)
}
