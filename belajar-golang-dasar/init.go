package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // jika ingin tetap ada tapi harus menambahkan "_"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
