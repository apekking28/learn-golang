package main

import "fmt"

/**
- Untuk method di rekomendasikan untuk menggunakan pointer
*/

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	king := Man{"king"}
	king.Married()

	fmt.Println(king.Name)
}
