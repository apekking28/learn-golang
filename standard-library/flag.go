package main

import (
	"flag"
	"fmt"
)

/**
 * Flag digunakan sebagai penanda kondisi tertentu
 * untuk mengontrol alur program.
 */

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "", "database host")
	var port *int = flag.Int("port", 0, "database port")

	// digunakan agar tidak mengembalikan value memori tapi real value nya
	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}
