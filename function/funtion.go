package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	// greet("Airell", 23)

// 	greet("Airell", "Jalan Sudirman")
// }

// // func greet(name string, age int8) {
// // 	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
// // }

// func greet(name, address string) {
// 	fmt.Println("Hello there! My name is Airell")
// 	fmt.Println("I live in", address)
// }

// Function (Return)
// func main() {
// 	var names = []string{"Airell", "Jordan"}
// 	var printMsg = greet("Heii", names)

// 	fmt.Println(printMsg)
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)
// 	return result
// }

// // Function (Returning Multiple Values)
// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)
// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference", circumference)
// }

// func calculate(d float64) (float64, float64) {
// 	// Menghitung Luas
// 	var area float64 = math.Pi * math.Pow(d/2, 2)
// 	// Menghitung Keliling
// 	var circumference = math.Pi * d
// 	return area, circumference
// }

// // Function (Predefined Return Value)
// func main() {
// 	var diameter float64 = 15
// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference", circumference)
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	// Menghitung Luas
// 	area = math.Pi * math.Pow(d/2, 2)
// 	// Menghitung Keliling
// 	circumference = math.Pi * d

// 	return
// }

// // Function (Variadic function #1)
// func main() {
// 	studentLists := print("Airell", "Nanda", "Mailo", "Marco")
// 	fmt.Printf("%v", studentLists)
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}

// 	return result
// }

// // Variadic Function #2
// func main() {
// 	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	result := sum(numberLists...)
// 	fmt.Println("Result:", result)
// }

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, v := range numbers {
// 		total += v
// 	}
// 	return total
// }

// Function (Variadic Function #3)
func main() {
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
