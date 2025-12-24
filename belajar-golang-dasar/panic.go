package main

import "fmt"

// panic akan mengentikan applikasi

func endApp() {
	fmt.Println("End application")
	// ini cara yang benar untuk recover karna kalo terjadi panic auto menghentikan aplikasi sehingga tidak menjalankan recovernya
	message := recover()
	fmt.Println("Terjadi error ,", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Upps error")
	}

	// cara salah menggunakan recover
	// message := recover()
	// fmt.Println("Terjadi error", message)
}

func main() {
	runApp(true)
	fmt.Println("kingek")
}
