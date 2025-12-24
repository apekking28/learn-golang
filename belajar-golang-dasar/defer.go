package main

import "fmt"

// defer tetap akan di eksekusi walaupun terjadi panic/error

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // defer akan memanggil function ini di akir eksekusi
	fmt.Println("Run application ...")
}

func main() {
	runApplication()
}
