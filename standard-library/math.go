package main

import (
	"fmt"
	"math"
)

/**
- 'math' package digunakan untuk operasi matematika

*/

func main() {
	fmt.Println(math.Ceil(1.40))  // membulatkan ke atas
	fmt.Println(math.Floor(1.60)) // membualatkan ke bawah
	fmt.Println(math.Round(1.60)) // membulatkan ke terdekat
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
