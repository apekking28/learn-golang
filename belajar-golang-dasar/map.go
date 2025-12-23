package main

import "fmt"

func main() {

	// cara 1

	// var person = map[string]string{}
	// person["name"] = "ilham"
	// person["address"] = "jakarta"

	person := map[string]string{
		"name":    "Ilham",
		"address": "Jakarta",
	}

	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["salah"]) // mengembalikan value sesuai dengan tipe datanya

	book := make(map[string]string)
	book["Title"] = "Buku golang"
	book["Author"] = "King"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
