package main

import "fmt"

/**
- Sturct adalah prototype data / template data
*/

type Customer struct {
	// kalo sama tipe nya
	// Name, Address string

	Name    string
	Address string
	Age     int
}

// method untuk struct sehingga menempel di struct ref nya
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var customer1 Customer

	fmt.Println(customer1)

	customer1.Name = "king"
	customer1.Address = "Jakarta"
	customer1.Age = 28

	fmt.Println(customer1)

	// cara akses value
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	// cara buat struct literal
	// cara 1
	joko := Customer{
		Name:    "joko",
		Address: "Banten",
		Age:     18,
	}

	// harus berurutan sesuai referensi struct nya
	budi := Customer{"Budi", "Amerika", 20}

	fmt.Println(joko)
	fmt.Println(budi)

	// cara memanggil method struct
	budi.sayHello("Budi")
}
