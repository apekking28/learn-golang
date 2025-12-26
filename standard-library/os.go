package main

import (
	"fmt"
	"os"
)

/**
 * Package os digunakan untuk berinteraksi dengan sistem operasi,
 * seperti mengakses environment variable, file system,
 * argument program, dan informasi proses.
 */

func main() {
	// 'os.Args' digunakan untuk mendapatkan argumen di cli ketika aplikasi sedang berjalan
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// 'os.Hostname()' digunakan untuk mendapatkan host
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
