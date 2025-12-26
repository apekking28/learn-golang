package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Package container/ring menyediakan struktur data circular linked list
// di mana elemen terakhir terhubung kembali ke elemen pertama.
// Ring mirip seperti jam dinding atau antrian melingkar,
// setelah elemen terakhir akan kembali ke elemen pertama.
// ring hanya bisa menyimpang value sesuai declarasi awal

func main() {
	// declarasi variabel ring
	var data *ring.Ring = ring.New(5)

	// cara manual
	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	// cara looping
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// cara untuk menampilkan data ring
	data.Do(func(val any) {
		fmt.Println(val)
	})

}
