package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Ilham")
	fmt.Println(result)

	fmt.Println(getHello("Uji"))
	fmt.Println(getHello("lana"))
}
