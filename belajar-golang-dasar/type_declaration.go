package main

import (
	"fmt"
	"reflect"
)

func main() {
	type NoKTP string

	var noktp NoKTP = "213138"

	var contoh string = "dasdada"

	contohNoKTP := NoKTP(contoh)

	fmt.Println(noktp)
	fmt.Println(contoh)
	fmt.Printf("Data type contohNoKTP %T", contohNoKTP)
	fmt.Println(reflect.TypeOf(noktp))
}
