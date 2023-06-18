package main

import "fmt"

func main1() {
	// Maps
	// name map[keyType]valueType
	//var m map[string]int
	// we can't assign a key value to map, because not point to address of memory
	// m["k1"] = 7
	//fmt.Println(m)

	//if m == nil {
	//m = make(map[string]int)
	//}
	//fmt.Println(m)
	//m["k1"] = 7
	//m["k2"] = 13
	//fmt.Println("map m:", m)

	//v1 := m["k1"]
	//fmt.Println("v1:", v1)

	//fmt.Println("m len:", len(m))

	//delete(m, "k2")
	//fmt.Println("map m:", m)

	//k2, ok := m["k2"]
	//fmt.Println("k2:", k2, "| ok:", ok)

	//_, ok := m["k2"]
	//fmt.Println("ok:", ok)

	n := map[string]int{"foo": 1, "bar": 2, "test": 3, "go": 4}
	fmt.Println("map n:", n)

	for key, value := range n {
		fmt.Println("key:", key, "| value:", value)
	}
}
