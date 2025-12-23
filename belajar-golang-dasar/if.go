package main

import "fmt"

func main() {
	name := "Uji anjay"

	if name == "Ilham" {
		fmt.Println("Hello Ilham")
	} else if name == "Lana" {
		fmt.Println("Hello Lana")
	} else if name == "Uji" {
		fmt.Println("Hello Uji")
	} else {
		fmt.Println("Hi, Boleh kenalan gk ?")
	}

	// length := len(name) -> ini cara yang umum di gunakan

	// if shorted kondisi
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
