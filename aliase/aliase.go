package main

import "fmt"

func main() {
	// // Deklarasi variable dengan tipe data uint8
	// var a uint8 = 10
	// var b byte // byte adalah alias dari uint8

	// b = a // no error, karena byte memiliki tipe data yang saman dengan uint8
	// _ = b

	// Mendeklarasi tipe data dengan nama second
	type second = uint
	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint
}
