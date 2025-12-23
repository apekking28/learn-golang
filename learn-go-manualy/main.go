package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var arr = [...]int{1, 2, 3}

	fmt.Println(arr)

	slice := []string{"ilham", "lana", "anajt"}

	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	fmt.Println(slice[0][2:5])
}
