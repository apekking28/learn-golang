package main

import "fmt"

/**
- type assertion digunakan untuk mengambil nilai konkret dari interface{} / any
- type assertion bisa ke tipe apa pun (int, string, struct, dll)
- jika tipe tidak sesuai, akan terjadi panic kecuali menggunakan idiom `value, ok`
*/

func random() any {
	return 100
}

func main() {
	var result any = random()
	// cara convert any ke string
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt)

	// check tipe nya
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unkwon value")
	}

}
