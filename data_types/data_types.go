package main

import "fmt"

func main() {
	// var first = 89
	// var second = -12

	// fmt.Printf("tipe data first : %T\n", first)
	// fmt.Printf("tipe data second : %T\n", second)

	// var first uint8 = 89
	// var second int8 = -12
	// fmt.Printf("tipe data first : %T\n", first)
	// fmt.Printf("tipe data second : %T\n", second)

	// var decimalNumber float32 = 3.63
	// fmt.Printf("decimal number: %f\n", decimalNumber)
	// fmt.Printf("decimal number: %.3f\n", decimalNumber)

	// var condition bool = true
	// fmt.Printf("is it permitted? %t \n", condition)

	// var message = "Halo"
	// fmt.Printf(message)
	// fmt.Println()
	// fmt.Printf("%s", message)

	var message = `Selamat datang di "Hacktiv8".
Salam Kenal :).
Mari belajar "Scalable Web Service with Go".`
	var message2 = `Selamat datang di "Hacktiv8".
	Salam Kenal :).
	Mari belajar "Scalable Web Service with Go".`

	fmt.Printf(message + "\n")
	fmt.Printf("%s"+"\n", message2)

	fmt.Println(2 ^ ^2)
}
