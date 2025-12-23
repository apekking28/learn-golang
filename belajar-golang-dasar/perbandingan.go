package main

import "fmt"

func main() {
	name1 := "king"
	name2 := "king"

	name3 := "ilham"

	fmt.Println(name1 == name2) // true
	fmt.Println(name1 == name3) // flase
	fmt.Println(name1 != name2) // false
}
