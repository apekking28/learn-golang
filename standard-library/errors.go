package main

import (
	"errors"
	"fmt"
)

var (
	ValidataionError = errors.New("Validation error")
	NotFoundError    = errors.New("Notfound error")
	UnknownError     = errors.New("Unknown error")
)

func GetById(id string) error {
	if id == "" {
		return ValidataionError
	} else if id != "king" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("king")

	// 'errors.Is()' cek type errornya
	if err != nil {
		if errors.Is(err, ValidataionError) {
			fmt.Println(err.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(err.Error())
		} else {
			fmt.Println(err.Error())
		}
	}
}
