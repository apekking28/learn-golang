package main

import "fmt"

func main() {
	a := 10
	b := 12
	c := a + b
	d := c + a
	e := d + b

	fmt.Println(e)
	// augmentent assisgment
	i := 10
	i += 15

	fmt.Println(i)
	// unneary operation
	j := 0

	j++ // j = j + 1
	j++ // j = j + 1

	fmt.Println(j)

	j-- // j = j - 1

	fmt.Println(j)
}
