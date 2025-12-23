package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are bloked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	// ini contoh anonymus function
	// cara 1
	blackList := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("King", blackList)

	// cara 2
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
