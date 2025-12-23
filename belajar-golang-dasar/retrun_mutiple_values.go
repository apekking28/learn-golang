package main

import "fmt"

func getFullName() (string, string) {
	return "Ilham", "Firmansyah"
}

func main() {

	// mengambil semua return functionnya
	// firstName, lastname := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastname)

	// mengambil salah satu nilai nya
	firstName, _ := getFullName()

	fmt.Println(firstName)

}
