package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa barat", "Indonesia"}
	// address2 := address1 // copy by value

	// address2.City = "jakarta"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah

	var address1 Address = Address{"Subang", "Jawa barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "jakarta"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah
}
