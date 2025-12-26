package main

import (
	"fmt"
	"strings"
)

/**
- 'strings' package digunakan untuk memanipulasi tipe data 'string'
*/

func main() {
	fmt.Println(strings.Contains("Ilham Firmansyah", "syah"))
	fmt.Println(strings.Split("Ilham Firmansyah", " "))
	fmt.Println(strings.ToLower("Ilham Firmansyah"))
	fmt.Println(strings.ToUpper("Ilham Firmansyah"))
	fmt.Println(strings.Trim("   Ilham Firmansyah   ", " "))             // untuk  membersihkan awal dan akhir string
	fmt.Println(strings.ReplaceAll("Ilham Firmansyah", "Ilham", "cina")) // mengganti string ke value yang diinginkan
}
