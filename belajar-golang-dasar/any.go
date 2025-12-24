package main

import "fmt"

/**
- any (alias dari interface{}) dapat menampung nilai dari tipe apa pun
- biasanya digunakan di layer paling umum (general / top-level abstraction)
- nilai bertipe any perlu type assertion atau type switch sebelum digunakan
*/

func Ups() any {
	// return 1
	return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
