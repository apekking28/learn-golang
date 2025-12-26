package main

import (
	"container/list"
	"fmt"
)

// Package container/list menyediakan struktur data doubly linked list.
// Setiap elemen saling terhubung ke elemen sebelum dan sesudahnya,
// mirip seperti rangkaian buku yang saling terhubung.
// struktur data seperti queue dan stack.

func main() {
	var data *list.List = list.New()

	data.PushBack("Ilham")
	data.PushBack("Firmansyah")
	data.PushBack("anjay")

	fmt.Println(data)

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Ilham

	next := head.Next() // Firmansyah
	fmt.Println(next.Value)

	next = head.Next() // anjay
	fmt.Println(next.Value)

	// cara looping list
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
