package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiakhir := nilaiAkhir > 80
	lulusNilaiAbsensi := absensi > 80

	lulus1 := lulusNilaiAbsensi && lulusNilaiakhir        // false
	tidakLulus := !(lulusNilaiAbsensi && lulusNilaiakhir) // true

	fmt.Println(lulus1)
	fmt.Println(tidakLulus)
}
