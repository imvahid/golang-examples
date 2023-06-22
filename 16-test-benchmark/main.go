package main

import (
	"fmt"
	"testing/ncode"
)

func main() {
	nID, err := ncode.NewNationalID("0110456785")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("nid is:", nID.IsValid)

	city, err := nID.GetCity()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("city is:", city)
}
