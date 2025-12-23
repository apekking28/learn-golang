package main

import "fmt"

func main() {
	var name string // declare variabel

	name = "Ilham Firmansyah" // assigned value variabel name
	fmt.Println(name)

	name = "apekking" // change value variabel name to 'apekking'
	fmt.Println(name)

	var address = "Jakarta" // go langsung tau tipe datanya dari value nya
	fmt.Println(address)

	age := "18" // cara paling simpel declarasi awal wariabel (hanya untuk pertama kali declarasi saja)
	fmt.Println(age)

	// Multiple variable
	var (
		firstName = "Ilham"
		lastName  = "Firmansyah"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
