package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }
	// fmt.Println("Selesai")

	// lebih sederhana
	for counter := 1; counter <= 10; counter++ { // for (init statment); (kondisi); (past statment) {}
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("Selesai")

	names := []string{"King", "Lana", "Ilham"}

	// cara biasa di pake
	for i := 0; i < len(names); i++ {
		fmt.Println("Index : ", i, "Value : ", names[i])
	}

	// for range
	// for range bisa digunakan di collection di go
	fmt.Println("Di bawah contoh for range")
	for key, value := range names {
		fmt.Println("key : ", key, "Value : ", value)
	}

	// jika tidak butuh key/value bisa di ganti dengan '_'
	fmt.Println("Di bawah contoh for range")
	fmt.Println("Kondisi key dihapus")
	for _, value := range names {
		fmt.Println(value)
	}

	fmt.Println("Kondisi value dihapus")
	for key, _ := range names {
		fmt.Println(key)
	}

}
