package main

import (
	"fmt"
	"sort"
)

// Package sort digunakan untuk mengurutkan koleksi data
// dengan mengimplementasikan sort.Interface atau menggunakan sort.Slice.

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	// long way

	// temp := u[i]
	// u[i] = u[j]
	// u[j] = temp

	// short way
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"king", 18},
		{"lana", 19},
		{"anjay", 17},
	}

	// convert ke userSlice
	// sort.Sort(UserSlice(users))

	// cara lebih singkat tanpa harus buat 'UserSlice' serta implementasi nya
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	fmt.Println(users)
}
