package main

import "fmt"

func getCompletedName() (firstName, midleName, lastName string) {
	firstName = "Ilham"
	midleName = " "
	lastName = "Firmansyah"
	return firstName, midleName, lastName
}

func main() {
	a, b, c := getCompletedName()
	fmt.Println(a, b, c)
}
