package main

import (
	"fmt"
)

func main() {
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Airell", "Nanda", "Mailo"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"
	// fmt.Printf("%#v\n", fruits)

	// Loop through elements
	// var fruits = [3]string{"apple", "banana", "mango"}
	// // cara pertama
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }
	// fmt.Println(strings.Repeat("#", 25))

	// // cara kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

	// Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
