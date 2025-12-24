package main

import "fmt"

/*
*
- interface itu kontrak yang dimana isi nya hanya method
- yang menggunakan implementasi harus ada method interface yang di gunakan
*/
type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"King"}
	SayHello(person)

	animal := Animal{"mewww"}
	SayHello(animal)
}
