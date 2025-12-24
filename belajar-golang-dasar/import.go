package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("king")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)           // tidak bisa di akses karna dari package ref
	// fmt.Println(helper.sayGoodBy("king")) // tidak bisa di akses karna dari package ref
}
