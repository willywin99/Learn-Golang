package main

import "fmt"

func main() {
	// var person map[string]string // Deklarasi
	// person = map[string]string{} // Inisialisasi
	// person["name"] = "Airell"
	// person["age"] = "23"
	// person["address"] = "Jalan Sudirman"

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])

	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jalan Sudirman",
	// }

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])

	// // Map (Looping with Map)
	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jalan Sudirman",
	// }

	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// // Map (Deleting Value)
	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jalan Sudirman",
	// }

	// fmt.Println("Before deleting", person)
	// delete(person, "age")
	// fmt.Println("After deleting", person)

	// Map (Detecting a Value)
	// var person = map[string]string{
	// 	"name":    "Airell",
	// 	"age":     "23",
	// 	"address": "Jalan Sudirman",
	// }

	// value, exist := person["age"]
	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value doesn't exist.")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value has been deleted.")
	// }

	// Combining Slice with Map
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
