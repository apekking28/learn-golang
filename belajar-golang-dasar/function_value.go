package main

import "fmt"

func goodBy(name string) string {
	return "good by " + name
}

func main() {

	// menyimpan func ke varibel
	contoh := goodBy
	misal := goodBy

	fmt.Println(contoh("king"))
	fmt.Println(misal("anjay"))
}
