package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "amerika"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah

	// membuat pointer baru dengan type Address
	// address2 = &Address{"Bandung", "Jawa barat", "Indonesia"}
	// fmt.Println(address1) // ikut berubah
	// fmt.Println(address2) // berubah

	*address2 = Address{"Bandung", "Jawa barat", "Indonesia"}

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah
}
