package main

import "fmt"

func main() {
	name := "lanaaajkgkjgkjgkjg"

	switch name {
	case "king":
		fmt.Println("Hello king")
	case "lana":
		fmt.Println("Hello lana")
	case "uji":
		fmt.Println("Hello uji")
	default:
		fmt.Println("Hi, Boleh kenalan ?")
	}

	// shorted case
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang benar")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
