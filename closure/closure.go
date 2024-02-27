package main

import "fmt"

// // Declare Closure in Variable
// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int

// 		for _, v := range numbers {
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}

// 	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
// 	fmt.Println(evenNumbers(numbers...))
// }

// Closure (IIFE)
func main() {
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")
	fmt.Println(isPalindrome)
}
