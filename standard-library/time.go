package main

import (
	"fmt"
	"time"
)

// Package time digunakan untuk mengelola dan memanipulasi waktu dan tanggal,
// termasuk timezone, parsing, dan formatting.

func main() {

	/*
		time.Now()   : mengambil waktu saat ini
		Local()      : mengonversi waktu ke timezone lokal
		time.Date()  : membuat waktu secara manual
		time.Parse() : mengubah string menjadi time.Time
		RFC3339      : format waktu standar ISO 8601
	*/

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	// ketika 'utc' di convert ke local (sesuai dengan waktu sistem) utc -> wib (+7 jam)
	fmt.Println(utc.Local())

	// parsing string ke waktu
	formatter := "2006-01-02 15:04:05"

	value := "2010-01-01 10:10:10"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
