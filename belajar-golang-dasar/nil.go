package main

import "fmt"

/**
- nil hanya bisa digunakan di type data interface, function, map, slice, pointer, chanel di luar ini tidak bisa
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("king")

	if data == nil {
		fmt.Println("data map masih koson")
	} else {
		fmt.Println(data["name"])
	}

	fmt.Println(data)
}
