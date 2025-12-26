package main

import (
	"fmt"
	"time"
)

// time.Duration adalah tipe data untuk merepresentasikan durasi waktu.
// Duration disimpan sebagai int64 dalam satuan nanodetik dan
// mendukung operasi matematika (penjumlahan, pengurangan, dll).

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 100 * time.Microsecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3)
}
