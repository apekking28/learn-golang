package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "king" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("a", nil)
	if err != nil {
		// terjadi error
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error :", validationErr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error :", notFoundError.Error())
		} else {
			fmt.Println("Unknow error:", err.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
