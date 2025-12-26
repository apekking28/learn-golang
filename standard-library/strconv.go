package main

import (
	"fmt"
	"strconv"
)

/**
- 'strconv' package digunakan untuk memrubah tipe data string ke tipe lain
*/

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(999)
	fmt.Println(stringInt)
}
