package main

import "fmt"

func main() {
	// tidak bisa di rubah jika constant

	// Multiple constant

	const (
		firstName string = "Apek"
		lastName         = "King"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
