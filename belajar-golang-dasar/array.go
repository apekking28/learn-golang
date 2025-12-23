package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "king"
	names[1] = "apek"
	names[2] = "sdasdad"

	fmt.Println(names)

	// hanya bisa digunakan ketika ingin mngisi value nya bukan pas declarasi
	var values = [...]int{
		90,
		80,
		70,
		100,
	}

	fmt.Println(values)

	fmt.Println(len(values))

	values[0] = 75
	fmt.Println(values)
}
