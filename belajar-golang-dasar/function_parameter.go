package main

import "fmt"

func sayHelloTo(firstName string, lastname string) {
	fmt.Println("Hello", firstName, lastname)
}

func main() {
	sayHelloTo("Ilham", "Firmansyah")
	sayHelloTo("Ahmad", "Fauzi")
	sayHelloTo("Maulana", "Rizky")
}
