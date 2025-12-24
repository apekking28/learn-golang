package helper

var version = "1.0.0"      // ini tidak bisa digunakan di package lain
var Application = "golang" // ini bisa digunakan di package lain

// ini tidak bisa digunakan di package lain
func sayGoodBy(name string) string {
	return "By " + name
}

// ini bisa digunakan di package lain
func SayHello(name string) string {
	return "Hello " + name
}
